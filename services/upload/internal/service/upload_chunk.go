package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"upload/internal/clients"
	"upload/internal/model"
	"upload/shared"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
)

/*
1. Resolve session
2. Validate status
3. Verify checksum
4. Write chunk to disk
5. INSERT chunk row (handle duplicate)
6. COUNT chunks
7. If complete:
  - update status → assembling
  - assemble
  - upload to S3
  - create file row
  - update status → completed

8. return nil
*/

// main func
func (s *UploadService) UploadChunk(ctx context.Context, input *UploadChunkInput) error {
	// 1,2
	session, err := s.mustAcceptChunks(ctx, input.UploadUUID)
	if err != nil {
		return err
	}

	// 3456
	err = s.handleChunk(ctx, session, input)
	if err != nil {
		if errors.Is(err, shared.ErrChunkAlreadyExists) {
			return err
		} else {
			return err
		}
	}
	//fmt.Println("chunk storeds")
	// 7. Count total chunks in chunks
	if complete, _ := s.isUploadComplete(ctx, session); !complete {
		return nil
	}

	// 1. Give the user an immediate "success" or "processing" signal
	// 2. Process the heavy lifting in the background
	go func() {
		// Use Background to ensure the process doesn't die when the request ends
		// But wrap it in a timeout so it doesn't hang forever
		bgCtx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)
		defer cancel()

		if err := s.finalizeUpload(bgCtx, session); err != nil {
			log.Printf("[ERROR] Finalization failed for session %d: %v", session.ID, err)
		}
		log.Println("[INFO] Upload complete for session", session.ID)
	}()

	return nil

}

func (s *UploadService) handleChunk(
	ctx context.Context,
	session *model.UploadSession,
	input *UploadChunkInput,
) error {

	// Create a hash calculator
	hasher := sha256.New()

	// TeeReader:
	// - data goes to hasher
	// - data continues downstream
	tee := io.TeeReader(input.ChunkBytes, hasher)

	// when bytes are being processed by tee, goes to two outlets at the same time
	sizeBytes, err := storeChunk(ctx, input.ChunkID, tee, session.ID)
	if err != nil {
		return err
	}

	// Verify checksum AFTER stream ends
	calculatedChecksum := hex.EncodeToString(hasher.Sum(nil))
	if input.CheckSum != "" && calculatedChecksum != input.CheckSum {
		return errors.New("checksum mismatch")
	}

	// Persist chunk metadata
	err = s.registry.Chunks.CreateChunk(ctx,
		&model.UploadChunk{
			SessionID:  session.ID,
			ChunkIndex: input.ChunkID,
			SizeBytes:  sizeBytes,
			CheckSum:   calculatedChecksum,
		})

	if errors.Is(err, shared.ErrChunkAlreadyExists) {
		return err
	}
	return err
}

func storeChunk(
	ctx context.Context,
	chunkID int,
	data io.Reader,
	sessionID int64,
) (int64, error) {

	path := chunkPath(sessionID, chunkID)

	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); err != nil {
		return 0, fmt.Errorf("session directory missing: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// io.Copy returns number of bytes written
	return io.Copy(f, data)
}

func chunkPath(sessionID int64, chunkID int) string {
	return filepath.Join(
		shared.UploadBasePath,
		strconv.FormatInt(sessionID, 10),
		fmt.Sprintf("%d.part", chunkID),
	)
}

func (s *UploadService) isUploadComplete(ctx context.Context, session *model.UploadSession) (bool, error) {
	count, err := s.registry.Chunks.CountBySession(ctx, session.ID)
	if err != nil {
		return false, err
	}
	return count == session.TotalChunks, nil
}

func (s *UploadService) finalizeUpload(ctx context.Context, session *model.UploadSession) error {
	s.registry.Sessions.UpdateSessionStatus(ctx, session.ID, "assembling")

	finalPath, err := s.assembleChunks(session.ID, session.TotalChunks)
	if err != nil {
		return s.fail(ctx, session.ID, err)
	}
	//	fmt.Println("Assenbled", finalPath)

	err = s.registry.Sessions.UpdateSessionStatus(ctx, session.ID, "uploading")
	if err != nil {
		return s.fail(ctx, session.ID, err)
	}
	// get mimeType of final file
	//	fileUUID := uuid.New()
	mimeType, err := DetectMimeFromFile(
		finalPath,
	)
	if err != nil {
		return s.fail(ctx, session.ID, err)
	}
	// Create File
	// calculate checksum of final file
	checksum, err := CalculateSHA256(finalPath)
	if err != nil {
		return s.fail(ctx, session.ID, err)
	}
	fileUUID := uuid.New()

	file := clients.CreateFileRequest{
		FileUUID:   fileUUID,
		UserID:     session.UserID,
		UploadUUID: session.UploadUUID,
		Name:       session.FileName,
		MimeType:   mimeType,
		SizeBytes:  session.FileSize,
		StorageKey: fmt.Sprintf(
			"%s%s/%s",
			shared.S3UsersPrefix,
			session.UserID,
			fileUUID,
		),
		CheckSum: checksum,
	}

	// detatch context and put
	backgroundCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 4*time.Minute)
	defer cancel()
	// upload to Cloud
	// send a cancel without context for handlning failed uploads
	err = s.storage.UploadFile(backgroundCtx, file.StorageKey, finalPath)
	if err != nil {
		return s.fail(ctx, session.ID, fmt.Errorf("s3 upload failed: %w", err))
	}

	// Create file row
	if err := s.fileClient.AddFile(backgroundCtx, &file); err != nil {
		log.Printf("[ERROR] failed to register file: %v, \n , FILE: %v", err, file)
		return s.fail(ctx, session.ID, fmt.Errorf("failed to register file: %w", err))
	}
	err = s.registry.Sessions.UpdateSessionStatus(ctx, session.ID, "completed")
	if err != nil {
		return err
	}
	err = s.removeSessionFolder(finalPath)
	if err != nil {
		return err
	}
	return nil

}

func (s *UploadService) removeSessionFolder(finalPath string) error {
	return os.RemoveAll(filepath.Dir(finalPath))
}

func (s *UploadService) mustAcceptChunks(ctx context.Context, id uuid.UUID) (*model.UploadSession, error) {
	session, err := s.registry.Sessions.GetSessionByUUID(ctx, id)
	if err != nil {
		return nil, err
	}
	if session.Status != "pending" {
		return nil, fmt.Errorf("session not accepting chunks")
	}
	return session, nil
}

func (s *UploadService) fail(ctx context.Context, sessionID int64, err error) error {
	log.Printf("[ERROR] Upload session %d failed: %v", sessionID, err)
	_ = s.registry.Sessions.UpdateSessionStatus(ctx, sessionID, "failed")

	sessionDir := filepath.Join(
		shared.UploadBasePath,
		strconv.FormatInt(sessionID, 10),
	)
	_ = os.RemoveAll(sessionDir)

	return err
}

func DetectMimeFromFile(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	mime, err := mimetype.DetectReader(f)
	if err != nil {
		return "", err
	}

	return mime.String(), nil
}

func CalculateSHA256(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hasher := sha256.New()

	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}

	sum := hasher.Sum(nil)
	return hex.EncodeToString(sum), nil
}
