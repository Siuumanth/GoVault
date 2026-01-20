package repository

import (
	"files/internal/model"
	"files/internal/service/files"

	"github.com/google/uuid"
)

type RepoRegistry struct {
	Metadata MetaDataRepository
	Files    FilesRepository
	//	Sharing   SharingRepository
	//	Shortcuts ShortcutsRepository
}

/*
/metadata
/files
/sharing
/shortcuts

type MetadataService interface {
	UpdateFileName(input *UpdateFileNameInput) error
	GetFileMetadata(fileID uuid.UUID, actorUserID uuid.UUID) (*model.File, error)
}

type FilesService interface {
	GetSingleFile(fileID uuid.UUID, actorUserID uuid.UUID) (*model.File, error)
	ListOwnedFiles(
		userID uuid.UUID,
		limit int,
		offset int,
	) ([]*model.File, error)
	ListSharedFiles(
		userID uuid.UUID,
		limit int,
		offset int,
	) ([]*model.File, error)
	MakeFileCopy(
		fileID uuid.UUID,
		actorUserID uuid.UUID,
	) (*model.File, error)
}

type SharingService interface {
	AddFileShares(input *AddFileSharesInput) error

	DeleteFileShare(
		fileID uuid.UUID,
		recipientUserID uuid.UUID,
	) error

	UpdateFileShare(input *UpdateFileShareInput) error

	ListFileShares(fileID uuid.UUID) ([]*model.FileShare, error)

	AddPublicAccess(fileID uuid.UUID) error
	RemovePublicAccess(fileID uuid.UUID) error
}

type ShortcutsService interface {
	CreateShortcut(
		fileID uuid.UUID,
		actorUserID uuid.UUID,
	) (*model.FileShortcut, error)
	DeleteShortcut(
		shortcutID uuid.UUID,
		actorUserID uuid.UUID,
	) error
}

*/

type MetaDataRepository interface {
	GetFileMetadata(fileID uuid.UUID) (*files.FileSummary, error)
	UpdateFileName(fileID uuid.UUID, newName string) (bool, error)
}

type FilesRepository interface {
	GetSingleFile(fileID uuid.UUID) (*files.FileSummary, error)
	ListOwnedFiles(userID uuid.UUID, limit int, offset int) ([]*files.FileSummary, error)
	ListSharedFiles(userID uuid.UUID, limit int, offset int) ([]*files.FileSummary, error)
	CreateFile(file *files.CreateFileParams) (*model.File, error)
}
