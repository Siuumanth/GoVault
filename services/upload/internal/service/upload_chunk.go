package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"upload/internal/model"
	"upload/shared"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
)

// TODO: Remove folder after complete
// TODO: update uploads by multipart or raw bytes or fail

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
	session, err := s.mustAcceptChunks(input.UploadUUID)
	if err != nil {
		return err
	}

	// 3456
	err = s.handleChunk(session, input)
	if err != nil {
		if errors.Is(err, shared.ErrChunkAlreadyExists) {
			return err
		} else {
			return err
		}
	}
	//fmt.Println("chunk storeds")
	// 7. Count total chunks in chunks
	if complete, _ := s.isUploadComplete(session); !complete {
		return nil
	}
	//	fmt.Println("finalising upload")
	// finalise upload by assemmblung, cloud
	err = s.finalizeUpload(ctx, session)
	if err != nil {
		return err
	}
	return nil

}

func (s *UploadService) handleChunk(session *model.UploadSession, input *UploadChunkInput) error {
	if err := verifyChecksum(input.ChunkBytes, input.CheckSum); err != nil {
		return err
	}

	if err := storeChunk(input.ChunkID, input.ChunkBytes, session.ID); err != nil {
		return err
	}

	err := s.registry.Chunks.CreateChunk(&model.UploadChunk{
		SessionID:  session.ID,
		ChunkIndex: input.ChunkID,
		SizeBytes:  int64(len(input.ChunkBytes)),
		CheckSum:   input.CheckSum,
	})
	if errors.Is(err, shared.ErrChunkAlreadyExists) {
		return err
	}
	return err
}

func (s *UploadService) isUploadComplete(session *model.UploadSession) (bool, error) {
	count, err := s.registry.Chunks.CountBySession(session.ID)
	if err != nil {
		return false, err
	}
	return count == session.TotalChunks, nil
}

func (s *UploadService) finalizeUpload(ctx context.Context, session *model.UploadSession) error {
	s.registry.Sessions.UpdateSessionStatus(session.ID, "assembling")

	finalPath, err := s.assembleChunks(session.ID, session.TotalChunks)
	if err != nil {
		return s.fail(session.ID, err)
	}
	//	fmt.Println("Assenbled", finalPath)

	err = s.registry.Sessions.UpdateSessionStatus(session.ID, "uploading")
	if err != nil {
		return s.fail(session.ID, err)
	}
	// get mimeType of final file
	fileUUID := uuid.New()
	mimeType, err := DetectMimeFromFile(
		finalPath,
	)
	if err != nil {
		return s.fail(session.ID, err)
	}
	// Create File
	// calculate checksum of final file
	checksum, err := CalculateSHA256(finalPath)
	if err != nil {
		return s.fail(session.ID, err)
	}

	file := model.File{
		FileUUID:  fileUUID,
		SessionID: session.ID,
		UserID:    session.UserID,
		Name:      session.FileName,
		SizeBytes: session.FileSize,
		MimeType:  mimeType,
		CheckSum:  &checksum,
		StorageKey: fmt.Sprintf(
			"%s%s/%s",
			shared.S3UsersPrefix,
			session.UserID,
			fileUUID,
		),
	}

	// upload to Cloud
	err = s.storage.UploadFile(
		ctx,
		file.StorageKey,
		finalPath,
	)
	if err != nil {
		return s.fail(session.ID, err)
	}

	// Create file row
	err = s.registry.Files.CreateFile(&file)
	if err != nil {
		return s.fail(session.ID, err)
	}

	err = s.registry.Sessions.UpdateSessionStatus(session.ID, "completed")
	if err != nil {
		return err
	}
	return nil

}

func (s *UploadService) mustAcceptChunks(id uuid.UUID) (*model.UploadSession, error) {
	session, err := s.registry.Sessions.GetSessionByUUID(id)
	if err != nil {
		return nil, err
	}
	if session.Status != "pending" {
		return nil, fmt.Errorf("session not accepting chunks")
	}
	return session, nil
}

func (s *UploadService) fail(sessionID int, err error) error {
	_ = s.registry.Sessions.UpdateSessionStatus(sessionID, "failed")
	return err
}

func storeChunk(chunkID int, chunk []byte, sessionID int) error {
	dirPath := filepath.Join(shared.UploadBasePath, fmt.Sprintf("%d", sessionID))

	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return err
	}

	savePath := filepath.Join(dirPath, fmt.Sprintf("%d.part", chunkID))

	err = os.WriteFile(savePath, chunk, 0644)
	if err != nil {
		return err
	}

	//fmt.Printf("chunk %d saved to %s\n", chunkID, savePath)
	return nil
}

// Verify that the checksum of the given data matches the expected string.
// If the checksum does not match, returns an error with the message "checksum mismatch".
// Otherwise, returns nil.
func verifyChecksum(data []byte, expected string) error {
	hash := sha256.Sum256(data)
	calculated := hex.EncodeToString(hash[:])

	if calculated != expected {
		return fmt.Errorf("checksum mismatch")
	}
	return nil
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
