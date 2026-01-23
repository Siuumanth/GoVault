package files

import (
	"context"
	"errors"
	"files/internal/model"
	"files/internal/repository"

	"github.com/google/uuid"
)

type FilesService struct {
	fileRepo  repository.FileRepository
	shareRepo repository.ShareRepository
}

func NewFilesService(f repository.FileRepository, s repository.ShareRepository) *FilesService {
	return &FilesService{
		fileRepo:  f,
		shareRepo: s,
	}
}

func (s *FilesService) UpdateFileName(
	ctx context.Context,
	fileID uuid.UUID,
	newName string,
) error {
	if newName == "" {
		return errors.New("file name cannot be empty")
	}

	// permission + existence check should live here or inside repo
	return s.fileRepo.UpdateFileName(ctx, fileID, newName)
}

func (s *FilesService) GetFileMetadata(
	ctx context.Context,
	fileID uuid.UUID,
	actorUserID uuid.UUID,
) (*model.FileSummary, error) {
	// repo handles joins + access check
	file, err := s.fileRepo.GetSingleFileData(ctx, fileID)
	if err != nil {
		return nil, err
	}

	return &model.FileSummary{
		FileUUID:  file.FileUUID,
		UserID:    file.UserID,
		Name:      file.Name,
		MimeType:  file.MimeType,
		SizeBytes: file.SizeBytes,
		CreatedAt: file.CreatedAt,
	}, nil
}
