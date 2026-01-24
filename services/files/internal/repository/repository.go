package repository

import (
	"context"
	"files/internal/model"

	"github.com/google/uuid"
)

type FileRepository interface {
	FetchFullFileByID(ctx context.Context, fileID uuid.UUID) (*model.File, error)
	FetchFileSummaryByID(ctx context.Context, fileID uuid.UUID) (*model.FileSummary, error)
	UpdateFileName(ctx context.Context, fileID uuid.UUID, newName string) (bool, error)
	FetchOwnedFiles(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*model.FileSummary, error)
	FetchSharedFiles(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*model.FileSummary, error)
	CreateFile(ctx context.Context, file *model.CreateFileParams) (*model.File, error)
	SoftDeleteFile(ctx context.Context, fileID uuid.UUID) error
	CheckFileOwnership(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) (bool, error)
}

type ShareRepository interface {
	CreateFileShare(ctx context.Context, p *model.FileShareParams) (*model.FileShare, error)
	FetchUserFileShare(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) (*model.FileShare, error)
	DeleteFileShare(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) error
	UpdateFileShare(ctx context.Context, p *model.FileShareParams) error
	FetchAllFileShares(ctx context.Context, fileID uuid.UUID) ([]*model.FileShare, error)
	IsFileSharedWithUser(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) (bool, error)
	IsFileEditableByUser(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) (bool, error)

	// Public Access Methods
	CreatePublicAccess(ctx context.Context, fileID uuid.UUID) error
	DeletePublicAccess(ctx context.Context, fileID uuid.UUID) error
	IsFilePublic(ctx context.Context, fileID uuid.UUID) (bool, error)
}

type ShortcutsRepository interface {
	CreateShortcut(ctx context.Context, fileUUID uuid.UUID, userID uuid.UUID) (*model.FileShortcut, error)
	DeleteShortcut(ctx context.Context, fileUUID uuid.UUID, userID uuid.UUID) error
}
