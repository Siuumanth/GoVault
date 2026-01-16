package service

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"upload/internal/model"
	"upload/shared"
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

func (s *UploadService) UploadChunk(input *UploadChunkInput) error {
	// 1.
	session, err := s.registry.Sessions.GetSessionByUUID(input.UploadUUID)

	if err != nil {
		return err
	}

	// 2.
	if session.Status != "uploading" {
		return fmt.Errorf("upload not in progress")
	}

	//4,5 compute checksum and check
	err = verifyChecksum(input.ChunkBytes, input.CheckSum)
	if err != nil {
		return err
	}

	// 3.
	err = storeChunk(input.ChunkID, input.ChunkBytes, session.ID)
	if err != nil {
		return err
	}

	// 6. Create chunk
	err = s.registry.Chunks.CreateChunk(&model.UploadChunk{
		SessionID:  session.ID,
		ChunkIndex: input.ChunkID,
		SizeBytes:  len(input.ChunkBytes),
		CheckSum:   input.CheckSum,
	})
	if err != nil {
		if errors.Is(err, shared.ErrChunkAlreadyExists) {
			return nil // idempotent success
		}
		return err
	}

	// 7. Count total chunks in chunks
	noChunksUploaded, err := s.registry.Chunks.GetSessionChunksCount(session.ID)
	if err != nil {
		return err
	}

	// 8 Count chunks
	if noChunksUploaded == session.TotalChunks {
		// Assemble and file upload left
		// TODO: Assemble file logic
		// TODO: Upload to S3 logic
	}
	return nil
}

func storeChunk(chunkID int, chunk []byte, sessionID int) error {
	savePath := filepath.Join(shared.UploadBasePath, fmt.Sprintf("%d/%d.part", sessionID, chunkID))

	fmt.Printf("Writing chunk %d to %s\n", chunkID, savePath)

	err := os.WriteFile(savePath, chunk, 0644)
	if err != nil {
		return err
	}
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
