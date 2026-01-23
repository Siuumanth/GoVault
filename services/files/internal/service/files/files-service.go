package files

import (
	"context"
	"errors"
	"files/internal/model"
	"files/internal/repository"
	"files/internal/service"

	"github.com/google/uuid"
)

/*
	type FilesService interface {
		UpdateFileName(ctx context.Context, in *UpdateFileNameInput) error
		GetSingleFile(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) (*model.FileSummary, error)

		ListOwnedFiles(ctx context.Context, in *ListOwnedFilesInput) ([]*model.FileSummary, error)
		ListSharedFiles(ctx context.Context, in *ListSharedFilesInput) ([]*model.FileSummary, error)

		MakeFileCopy(ctx context.Context, in *MakeFileCopyInput) (*model.File, error)
	}
*/
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

func (s *FilesService) UpdateFileName(ctx context.Context, in *service.UpdateFileNameInput) error {
	if in.NewName == "" {
		return errors.New("file name cannot be empty")
	}

	// permission + existence check should live here or inside repo
	deleted, err := s.fileRepo.UpdateFileName(ctx, in.FileID, in.NewName)

	if !deleted {
		return errors.New("file not found or soft-deleted")
	}
	return err

}

// specific input type fot his
func (s *FilesService) GetFileMetadata(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) (*model.FileSummary, error) {
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

func (s *FilesService) ListOwnedFiles(ctx context.Context, in *service.ListOwnedFilesInput) ([]*model.FileSummary, error) {
	// definition
	var files []*model.FileSummary
	// repo handles joins + access check
	files, err := s.fileRepo.FetchOwnedFiles(ctx, in.UserID, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (s *FilesService) ListSharedFiles(ctx context.Context, in *service.ListSharedFilesInput) ([]*model.FileSummary, error) {

}
