package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type DownloadResponse struct {
	DownloadURL string `json:"download_url"`
}

func (s *DownloadService) GetDownloadURL(
	ctx context.Context,
	userID uuid.UUID,
	fileID uuid.UUID,
) (*DownloadResponse, error) {

	file, err := s.repos.Files.GetByID(fileID)
	if err != nil {
		return nil, err
	}

	if file.UserID != userID {
		return nil, errors.New("unauthorized")
	}

	url, err := s.storage.GenerateDownloadURL(
		ctx,
		file.StorageKey,
		300, // 5 minutes
	)
	if err != nil {
		return nil, err
	}

	return &DownloadResponse{
		DownloadURL: url,
	}, nil
}
