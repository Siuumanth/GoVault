package service

import (
	"context"
	"files/internal/model"
	"files/internal/repository"
	"files/internal/service/inputs"
	"files/internal/storage"

	"github.com/google/uuid"
)

type ServiceRegistry struct {
	Files     FilesService
	Sharing   SharesService
	Shortcuts ShortcutsService
}

// func that takes in repo registry and returns service registry:
func NewServiceRegistry(r *repository.RepoRegistry, storage storage.FileStorage) *ServiceRegistry {
	return &ServiceRegistry{
		Files:     NewFilesService(r.Files, r.Shares, storage),
		Sharing:   NewSharesService(registry),
		Shortcuts: NewShortcutsService(registry),
	}
}

type FilesService interface {
	UpdateFileName(ctx context.Context, in *inputs.UpdateFileNameInput) error
	GetSingleFileSummary(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) (*model.FileSummary, error)

	ListOwnedFiles(ctx context.Context, in *inputs.ListOwnedFilesInput) ([]*model.FileSummary, error)
	ListSharedFiles(ctx context.Context, in *inputs.ListSharedFilesInput) ([]*model.FileSummary, error)

	MakeFileCopy(ctx context.Context, in *inputs.MakeFileCopyInput) (*model.File, error)
	SoftDeleteFile(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) error
}
type SharesService interface {
	AddFileShares(ctx context.Context, in *inputs.AddFileSharesInput) error
	UpdateFileShare(ctx context.Context, in *inputs.UpdateFileShareInput) error
	RemoveFileShare(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID, recipientUserID uuid.UUID) error
	ListFileShares(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) ([]*model.FileShare, error)
	AddPublicAccess(ctx context.Context, in *inputs.AddPublicAccessInput) error
	RemovePublicAccess(ctx context.Context, in *inputs.RemovePublicAccessInput) error
}

type ShortcutsService interface {
	CreateShortcut(ctx context.Context, in *inputs.CreateShortcutInput) (*model.FileShortcut, error)
	DeleteShortcut(ctx context.Context, in *inputs.DeleteShortcutInput) error
}
