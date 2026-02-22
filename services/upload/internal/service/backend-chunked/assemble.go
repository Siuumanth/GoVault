package backend

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"upload/shared"
)

func (s *ProxyUploadService) assembleChunks(sessionID int64, totalChunks int64) (string, error) {
	sessionDir := filepath.Join(shared.UploadBasePath, fmt.Sprintf("%d", sessionID))
	finalPath := filepath.Join(sessionDir, "final")

	out, err := os.Create(finalPath)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// copy all individual chunks into the final file
	for i := 0; i < int(totalChunks); i++ {
		partPath := filepath.Join(sessionDir, fmt.Sprintf("%d.part", i))

		in, err := os.Open(partPath)
		if err != nil {
			return "", err
		}

		_, err = io.Copy(out, in)
		in.Close()
		if err != nil {
			return "", err
		}
	}

	return finalPath, nil
}
