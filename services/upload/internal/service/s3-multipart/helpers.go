package multipart

import (
	"context"
	"log"
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"upload/shared"
)

func calculateTotalParts(fileSize int64, partSize int64) int64 {
	if partSize <= 0 {
		return 0
	}
	// Standard "Ceil" logic: (a + b - 1) / b
	return (fileSize + partSize - 1) / partSize
}

func getMimeType(filename string) string {
	mimeType := mime.TypeByExtension(filepath.Ext(filename))
	if mimeType == "" {
		mimeType = "application/octet-stream"
	}
	return mimeType
}

func (s *MultipartUploadService) fail(ctx context.Context, sessionID int64, err error) error {
	log.Printf("[ERROR] Upload session %d failed: %v", sessionID, err)
	_ = s.registry.Sessions.UpdateSessionStatus(ctx, sessionID, "failed")

	sessionDir := filepath.Join(
		shared.UploadBasePath,
		strconv.FormatInt(sessionID, 10),
	)
	_ = os.RemoveAll(sessionDir)

	return err
}
