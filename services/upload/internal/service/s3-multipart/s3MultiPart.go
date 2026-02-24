package multipart

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
	"upload/internal/clients"
	"upload/internal/model"
	"upload/internal/service/inputs"
	"upload/shared"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
)

// type ServiceMethods interface {
// 	UploadSession(ctx context.Context, inputs *UploadSessionInput) (*model.UploadSession, error)
// 	UploadChunk(ctx context.Context, inputs *UploadChunkInput) (*model.UploadChunk, error)
// 	GetUploadStatus(ctx context.Context, upload_uuid uuid.UUID) (*model.UploadSession, error)
// }

func (s *MultipartUploadService) UploadSession(ctx context.Context, in *inputs.UploadSessionInput) (*model.UploadSession, error) {
	// form model
	var session model.UploadSession
	fileUUID := uuid.New()
	n := calculateTotalParts(in.FileSizeBytes, in.PartSize)
	session.UploadUUID = fileUUID
	session.FileName = in.FileName
	session.FileSize = in.FileSizeBytes
	session.UserID = in.UserID
	session.TotalParts = n
	session.UploadMethod = "multipart"
	// 2. Generate S3 Key (e.g., users/UUID/filename)
	objectKey := fmt.Sprintf(
		"%s%s/%s",
		shared.S3UsersPrefix,
		session.UserID,
		fileUUID.String(),
	)

	// 3. Talk to S3 to get the UploadID
	uploadID, err := s.storage.InitiateMultipart(ctx, objectKey)
	if err != nil {
		return nil, fmt.Errorf("failed to initiate s3 multipart: %w", err)
	}
	session.StorageUploadID = &uploadID

	// 4. Save to Database
	_, err = s.registry.Sessions.CreateSession(ctx, &session)
	if err != nil {
		// If DB fails, abort S3 upload to prevent orphan sessions and storage costs
		log.Printf("[CLEANUP] DB failed, aborting S3 multipart for key: %s", objectKey)

		// We use a background context or a copy of the context to ensure
		// the abort attempt finishes even if the original request times out
		abortErr := s.storage.AbortMultipart(context.Background(), objectKey, uploadID)
		if abortErr != nil {
			log.Printf("[ERROR] Critical: Failed to abort S3 multipart after DB error: %v", abortErr)
		}

		return nil, fmt.Errorf("failed to create session in DB: %w", err)
	}

	return &session, nil
}

// AddS3Part records the ETag received from the frontend after a direct S3 upload
func (s *MultipartUploadService) AddS3Part(ctx context.Context, uploadUUID uuid.UUID, actorID uuid.UUID, input *inputs.AddPartInput) error {
	// 1. resolve Session and validate status
	session, err := s.registry.Sessions.GetSessionByUUID(ctx, uploadUUID)
	if err != nil {
		return err
	}

	if session.UserID != actorID {
		return errors.New("Unauthorized")
	}

	if session.UploadMethod != "multipart" {
		log.Println("session IS : ", session)
		return errors.New("invalid upload method: session is not multipart")
	}

	if session.Status != "pending" && session.Status != "uploading" {
		return errors.New("session not accepting parts")
	}

	// 2. If this is the first part, move status to 'uploading'
	if session.Status == "pending" {
		_ = s.registry.Sessions.UpdateSessionStatus(ctx, session.ID, "uploading")
	}

	// 3. Persist the part metadata
	// Using the internal session.ID (int64) as the foreign key
	part := &model.UploadPart{
		SessionID:  session.ID,
		PartNumber: input.PartNumber,
		SizeBytes:  input.SizeBytes,
		Etag:       input.Etag,
	}

	err = s.registry.Parts.CreatePart(ctx, part)
	if err != nil {
		if errors.Is(err, shared.ErrPartAlreadyExists) {
			return err
		}
		return fmt.Errorf("failed to store part metadata: %w", err)
	}

	// check ing if all part uppoaded wil be init by frontned

	return nil
}

// decompose function
func (s *MultipartUploadService) CompleteS3Multipart(ctx context.Context, uploadUUID uuid.UUID, actorID uuid.UUID) error {
	// 1. Get session
	session, err := s.registry.Sessions.GetSessionByUUID(ctx, uploadUUID)
	if err != nil {
		return err
	}
	// check ownership
	if session.UserID != actorID {
		return errors.New("Unauthorized")
	}
	// RECONSTRUCT the exact same key used in InitiateMultipart
	storageKey := fmt.Sprintf("%s%s/%s", shared.S3UsersPrefix, session.UserID, session.UploadUUID.String())

	if session.UploadMethod != "multipart" || session.StorageUploadID == nil {
		return errors.New("invalid upload method or missing S3 session")
	}

	// 2. Fetch parts from DB
	parts, err := s.registry.Parts.GetPartsBySession(ctx, session.ID)
	if err != nil {
		return err
	}

	// 3. Check if all parts have been uploaded
	if len(parts) != int(session.TotalParts) {
		return fmt.Errorf("missing parts: have %d, want %d", len(parts), session.TotalParts)
	}

	// 4. Assemble parts request for AWS
	var completedParts []types.CompletedPart
	for _, p := range parts {
		completedParts = append(completedParts, types.CompletedPart{
			ETag:       aws.String(p.Etag),
			PartNumber: aws.Int32(int32(p.PartNumber)),
		})
	}

	// Detach context for finalization
	bgCtx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()

	// FIX: Capturing both the location string and the error
	location, err := s.storage.CompleteMultipart(bgCtx, storageKey, *session.StorageUploadID, completedParts)
	if err != nil {
		return s.fail(ctx, session.ID, fmt.Errorf("s3 assembly failed: %w", err))
	}

	// Optional: Log the final S3 location for debugging
	log.Printf("[INFO] S3 assembly complete: %s", location)

	// 5. detect MimeType
	mimeType := getMimeType(session.FileName)

	// 6. Request file client to save file
	err = s.fileClient.AddFile(bgCtx, &clients.CreateFileRequest{
		FileUUID:   session.UploadUUID,
		UserID:     session.UserID,
		UploadUUID: session.UploadUUID,
		Name:       session.FileName,
		SizeBytes:  session.FileSize,
		MimeType:   mimeType,
		StorageKey: storageKey,
		CheckSum:   "s3-verified",
	})

	if err != nil {
		return s.fail(ctx, session.ID, err)
	}

	// 7. Mark as completed
	return s.registry.Sessions.UpdateSessionStatus(ctx, session.ID, "completed")
}
func (s *MultipartUploadService) GenerateAllPartURLs(
	ctx context.Context,
	uploadUUID uuid.UUID,
) ([]inputs.PresignedPart, error) {

	session, err := s.registry.Sessions.GetSessionByUUID(ctx, uploadUUID)
	if err != nil {
		return nil, err
	}

	objectKey := fmt.Sprintf(
		"%s%s/%s",
		shared.S3UsersPrefix,
		session.UserID,
		session.UploadUUID.String(),
	)

	var parts []inputs.PresignedPart

	for i := 1; i <= int(session.TotalParts); i++ {

		url, err := s.storage.PresignUploadPart(
			ctx,
			objectKey,
			*session.StorageUploadID,
			int32(i),
			15*time.Minute,
		)
		if err != nil {
			return nil, err
		}

		parts = append(parts, inputs.PresignedPart{
			PartNumber: i,
			URL:        url,
		})
	}

	return parts, nil
}
