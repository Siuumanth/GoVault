package multipart

import (
	"context"
	"fmt"
	"upload/internal/model"
	"upload/internal/service/inputs"
	"upload/shared"

	"github.com/google/uuid"
)

// type ServiceMethods interface {
// 	UploadSession(ctx context.Context, inputs *UploadSessionInput) (*model.UploadSession, error)
// 	UploadChunk(ctx context.Context, inputs *UploadChunkInput) (*model.UploadChunk, error)
// 	GetUploadStatus(ctx context.Context, upload_uuid uuid.UUID) (*model.UploadSession, error)
// }

func (s *UploadService) UploadSession(ctx context.Context, in *inputs.UploadSessionInput) (*model.UploadSession, error) {
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

	_, err = s.registry.Sessions.CreateSession(ctx, &session)
	if err != nil {
		// TODO: If DB fails, you should technically call s.storage.AbortMultipart
		return nil, err
	}

	return &session, nil
}
