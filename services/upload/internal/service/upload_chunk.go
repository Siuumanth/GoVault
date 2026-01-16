package service

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"upload/internal/model"
	"upload/shared"
)

/*
1. Resolve session by upload_uuid
2. If status != uploading → reject
3. Stream bytes → temp file
4. Compute checksum while streaming
5. If checksum mismatch → reject
6. INSERT upload_chunks (unique constraint)
7. COUNT(upload_chunks)
8. If count == total_chunks:
       - set status = assembling
       - assemble file
       - upload to S3
       - create files row
       - set status = completed
9. Return success
*/

func (s *UploadService) UploadChunk(input *UploadChunkInput) error {
	// 1.
	chunk, err := s.registry.Sessions.GetSessionByUUID(input.UploadUUID)

	if err != nil {
		return err
	}

	// 2.
	if chunk.Status != "uploading" {
		return fmt.Errorf("upload not in progress")
	}

	// 3.
	err = storeChunk(input.ChunkID, input.ChunkBytes, chunk.ID)
	if err != nil {
		return err
	}
	//4,5 compute checksum and check
	err = verifyChecksum(input.ChunkBytes, input.CheckSum)
	if err != nil {
		return err
	}

	// 6.
	err = s.registry.Chunks.CreateChunk(&model.UploadChunk{
		SessionID:  chunk.ID,
		ChunkIndex: input.ChunkID,
		SizeBytes:  len(input.ChunkBytes),
		CheckSum:   input.CheckSum,
	})
	if err != nil {
		return err
	}

}

func storeChunk(chunkID int, chunk []byte, sessionID int) error {
	savePath := filepath.Join(shared.UploadBasePath, fmt.Sprintf("%d/%d", sessionID, chunkID))

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
