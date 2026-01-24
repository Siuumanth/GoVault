package service

import (
	"context"
	"files/internal/model"

	"github.com/google/uuid"
)

type ServiceRegistry struct {
	Files     FilesService
	Sharing   SharingService
	Shortcuts ShortcutsService
}

type FilesService interface {
	UpdateFileName(ctx context.Context, in *UpdateFileNameInput) error
	GetSingleFileSummary(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) (*model.FileSummary, error)

	ListOwnedFiles(ctx context.Context, in *ListOwnedFilesInput) ([]*model.FileSummary, error)
	ListSharedFiles(ctx context.Context, in *ListSharedFilesInput) ([]*model.FileSummary, error)

	MakeFileCopy(ctx context.Context, in *MakeFileCopyInput) (*model.File, error)
	SoftDeleteFile(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) error

	// helper
	// checkFileAccess(
	// 	ctx context.Context,
	// 	fileID uuid.UUID,
	// 	actorUserID uuid.UUID,
	// ) (*model.File, error)
	// isUserOwnerOfFile(ctx context.Context, fileID *uuid.UUID, userID *uuid.UUID) (bool, error)
}
type SharingService interface {
	AddFileShares(ctx context.Context, in *AddFileSharesInput) error
	UpdateFileShare(ctx context.Context, in *UpdateFileShareInput) error
	RemoveFileShare(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID, recipientUserID uuid.UUID) error
	ListFileShares(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) ([]*model.FileShare, error)
	AddPublicAccess(ctx context.Context, in *AddPublicAccessInput) error
	RemovePublicAccess(ctx context.Context, in *RemovePublicAccessInput) error

	// // helper
	// doesUserHaveEditPermissions(ctx context.Context, fileID *uuid.UUID, userID *uuid.UUID) (bool, error)
}

type ShortcutsService interface {
	CreateShortcut(ctx context.Context, in *CreateShortcutInput) (*model.FileShortcut, error)
	DeleteShortcut(ctx context.Context, in *DeleteShortcutInput) error
}
