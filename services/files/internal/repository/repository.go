package repository

import (
	"context"
	"files/internal/model"

	"github.com/google/uuid"
)

type FileRepository interface {
	UpdateFileName(ctx context.Context, fileID uuid.UUID, newName string) (bool, error)
	GetSingleFileData(ctx context.Context, fileID uuid.UUID) (*model.FileSummary, error)
	FetchOwnedFiles(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*model.FileSummary, error)
	FetchSharedFiles(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*model.FileSummary, error)
	CreateFile(ctx context.Context, file *model.CreateFileParams) (*model.File, error)
}

type ShareRepository interface {
	CreateFileShare(ctx context.Context, p *model.FileShareParams) (*model.FileShare, error)
	DeleteFileShare(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) error
	UpdateFileShare(ctx context.Context, p *model.FileShareParams) error
	FetchFileShares(ctx context.Context, fileID uuid.UUID) ([]*model.FileShare, error)
	CreatePublicAccess(ctx context.Context, fileID uuid.UUID) error
	DeletePublicAccess(ctx context.Context, fileID uuid.UUID) error
}

type ShortcutsRepository interface {
	CreateShortcut(ctx context.Context, fileUUID uuid.UUID, userID uuid.UUID) (*model.FileShortcut, error)
	DeleteShortcut(ctx context.Context, fileUUID uuid.UUID, userID uuid.UUID) error
}
