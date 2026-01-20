package service

import (
	"files/internal/model"

	"github.com/google/uuid"
)

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
