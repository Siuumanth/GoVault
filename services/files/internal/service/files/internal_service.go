package files

import (
	"context"
	"files/internal/model"
	"files/internal/shared"
	"time"

	"github.com/google/uuid"
)

// GetMetadata returns the full model (StorageKey, Checksum, etc.)
// used by Download/Upload services.
func (s *FileService) GetDownloadDetails(ctx context.Context, fileID uuid.UUID) (*model.DownloadResponse, error) {
	d, err := s.fileRepo.FetchDownloadInfo(ctx, fileID)
	if err != nil {
		return nil, err
	}
	expiresAt := time.Now().Add(shared.DOWNLOAD_LINK_TTL)

	return &model.DownloadResponse{StorageKey: d.StorageKey, ExpiresAt: expiresAt}, nil
}

// SystemVerifyAccess reuses your existing Public Service logic
// to ensure the user has permission to the file before a download starts.
func (s *FileService) SystemVerifyAccess(ctx context.Context, fileID, userID uuid.UUID) (bool, error) {
	_, err := s.checkFileAccess(ctx, fileID, userID)
	if err != nil {
		return false, err
	}
	return true, nil
}
