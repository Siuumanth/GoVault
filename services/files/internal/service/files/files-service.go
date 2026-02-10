package files

import (
	"context"
	"files/internal/model"
	"files/internal/repository"
	"files/internal/service/inputs"
	"files/internal/shared"
	"files/internal/storage"
	"fmt"

	"github.com/google/uuid"
)

/*
	type FilesService interface {
		UpdateFileName(ctx context.Context, in *UpdateFileNameInput) error
		GetSingleFileSummary(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) (*model.FileSummary, error)

		ListOwnedFiles(ctx context.Context, in *ListOwnedFilesInput) ([]*model.FileSummary, error)
		ListSharedFiles(ctx context.Context, in *ListSharedFilesInput) ([]*model.FileSummary, error)

		MakeFileCopy(ctx context.Context, in *MakeFileCopyInput) (*model.File, error)
		SoftDeleteFile(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) error
	}
*/
type FileService struct {
	fileRepo  repository.FilesRepository
	shareRepo repository.SharesRepository
	storage   storage.FileStorage
}

func NewFileService(f repository.FilesRepository, s repository.SharesRepository, fs storage.FileStorage) *FileService {
	return &FileService{
		fileRepo:  f,
		shareRepo: s,
		storage:   fs,
	}
}

func (s *FileService) UpdateFileName(ctx context.Context, in *inputs.UpdateFileNameInput) error {

	// if file is owned or user is editor then only allow
	canEdit, err := s.canUserEditFile(ctx, in.FileID, in.ActorUserID)
	if err != nil {
		return err
	}
	if !canEdit {
		return shared.ErrUnauthorized
	}

	// permission + existence check should live here or inside repo
	success, err := s.fileRepo.UpdateFileName(ctx, in.FileID, in.NewName)

	if !success {
		return shared.ErrRowNotFound
	}
	return err

}

func (s *FileService) GetSingleFileSummary(ctx context.Context, fileID uuid.UUID, actorUserID *uuid.UUID) (*model.FileSummary, error) {
	/*
		First check if user is owner of file
		else check if file is public
		else check if user has access by shared
	*/

	isAllowed, err := s.checkFileAccess(ctx, fileID, actorUserID)

	if err != nil {
		return nil, err
	} else if !isAllowed {
		return nil, shared.ErrUnauthorized
	}

	file, err := s.fileRepo.FetchFileSummaryByID(ctx, fileID)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *FileService) ListOwnedFiles(ctx context.Context, in *inputs.ListOwnedFilesInput) ([]*model.FileSummary, error) {
	// definition
	var files []*model.FileSummary
	// repo handles joins + access check
	files, err := s.fileRepo.FetchOwnedFiles(ctx, in.UserID, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (s *FileService) ListSharedFiles(ctx context.Context, in *inputs.ListSharedFilesInput) ([]*model.FileSummary, error) {
	// definition
	var files []*model.FileSummary
	// repo handles joins + access check
	files, err := s.fileRepo.FetchSharedFiles(ctx, in.UserID, in.Limit, in.Offset)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (s *FileService) MakeFileCopy(ctx context.Context, in *inputs.MakeFileCopyInput) (*model.File, error) {

	// first check if user has access to the file
	isAllowed, err := s.checkFileAccess(ctx, in.FileID, &in.ActorUserID)
	if err != nil {
		return nil, err
	} else if !isAllowed {
		return nil, shared.ErrUnauthorized
	}

	src, err := s.fileRepo.FetchFullFileByID(ctx, in.FileID)
	if err != nil {
		return nil, err
	}

	newUUID := uuid.New()

	storageKey := fmt.Sprintf(
		"%s%s/%s",
		shared.S3UsersPrefix,
		in.ActorUserID.String(),
		newUUID.String(),
	)
	// first store then add to db
	if err := s.storage.Copy(ctx, src.StorageKey, storageKey); err != nil {
		return nil, err
	}

	params := &model.CreateFileParams{
		UploadUUID: nil, // lets keep it 0 for copied files
		FileUUID:   newUUID,
		UserID:     in.ActorUserID,
		Name:       src.FileName,
		MimeType:   src.MimeType,
		SizeBytes:  src.SizeBytes,
		Checksum:   src.Checksum,
		StorageKey: storageKey,
	}

	newFile, err := s.fileRepo.CreateFile(ctx, params)
	if err != nil {
		return nil, err
	}

	return newFile, nil
}

func (s *FileService) SoftDeleteFile(ctx context.Context, fileID uuid.UUID, actorUserID *uuid.UUID) error {

	// first check if user has access to the file
	file, err := s.fileRepo.FetchFileSummaryByID(ctx, fileID)
	if err != nil {
		return err
	}

	if file.UserID != *actorUserID {
		return shared.ErrUnauthorized
	}

	return s.fileRepo.SoftDeleteFile(ctx, fileID)
}
