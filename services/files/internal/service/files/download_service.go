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
func (s *FileService) GetDownloadDetails(
	ctx context.Context,
	fileID uuid.UUID,
	userID *uuid.UUID,
) (*model.DownloadResponse, error) {

	// Verify access
	isAllowed, err := s.checkFileAccess(ctx, fileID, userID)
	if err != nil {
		return nil, err
	}
	if !isAllowed {
		return nil, shared.ErrUnauthorized
	}

	// Fetch file metadata
	d, err := s.fileRepo.FetchDownloadInfo(ctx, fileID)
	if err != nil {
		return nil, err
	}

	// Generate presigned URL
	downloadURL, err := s.storage.GenerateDownloadURL(
		ctx,
		d.StorageKey,
		shared.DOWNLOAD_LINK_TTL,
	)
	if err != nil {
		return nil, err
	}

	// Return actual download URL
	return &model.DownloadResponse{
		DownloadURL: downloadURL,
		ExpiresAt:   time.Now().Add(shared.DOWNLOAD_LINK_TTL),
		FileName:    d.FileName,
	}, nil
}
