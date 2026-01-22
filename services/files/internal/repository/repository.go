package repository

import (
	"context"
	"files/internal/model"
	"files/internal/service/files"
	"files/internal/service/share"

	"github.com/google/uuid"
)

type RepoRegistry struct {
	Metadata  MetaDataRepository
	File      FileRepository
	Sharing   ShareRepository
	Shortcuts ShortcutsRepository
}

type MetaDataRepository interface {
	GetFileMetadata(ctx context.Context, fileID uuid.UUID) (*files.FileSummary, error)
	UpdateFileName(ctx context.Context, fileID uuid.UUID, newName string) (bool, error)
}

type FileRepository interface {
	GetSingleFile(ctx context.Context, fileID uuid.UUID) (*files.FileSummary, error)
	ListOwnedFiles(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*files.FileSummary, error)
	ListSharedFiles(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*files.FileSummary, error)
	CreateFile(ctx context.Context, file *files.CreateFileParams) (*model.File, error)
}

type ShareRepository interface {
	CreateFileShare(ctx context.Context, p *share.FileShareParams) (*model.FileShare, error)
	DeleteFileShare(ctx context.Context, fileID uuid.UUID, userID uuid.UUID) error
	UpdateFileShare(ctx context.Context, p *share.FileShareParams) error
	FetchFileShares(ctx context.Context, fileID uuid.UUID) ([]*model.FileShare, error)
	CreatePublicAccess(ctx context.Context, fileID uuid.UUID) error
	DeletePublicAccess(ctx context.Context, fileID uuid.UUID) error
}

type ShortcutsRepository interface {
	CreateShortcut(ctx context.Context, fileUUID uuid.UUID, userID uuid.UUID) (*model.FileShortcut, error)
	DeleteShortcut(ctx context.Context, fileUUID uuid.UUID, userID uuid.UUID) error
}
