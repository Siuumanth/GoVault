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
func (s *FileService) GetDownloadDetails(ctx context.Context, fileID uuid.UUID, userID *uuid.UUID) (*model.DownloadResponse, error) {
	// first verify if user can acccess file
	isAllowed, err := s.checkFileAccess(ctx, fileID, userID)
	if err != nil {
		return nil, err
	} else if !isAllowed {
		return nil, shared.ErrUnauthorized
	}

	d, err := s.fileRepo.FetchDownloadInfo(ctx, fileID)
	if err != nil {
		return nil, err
	}
	expiresAt := time.Now().Add(shared.DOWNLOAD_LINK_TTL)

	return &model.DownloadResponse{StorageKey: d.StorageKey, ExpiresAt: expiresAt}, nil
}
