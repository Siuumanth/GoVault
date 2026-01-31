package service

import (
	"context"
	"files/internal/model"
	"files/internal/repository"
	"files/internal/service/files"
	"files/internal/service/inputs"
	"files/internal/service/shares"
	"files/internal/service/shortcuts"
	"files/internal/storage"

	"github.com/google/uuid"
)

type ServiceRegistry struct {
	Files     *files.FileService
	Shares    *shares.ShareService
	Shortcuts *shortcuts.ShortcutService
}

// func that takes in repo registry and returns service registry:
func NewServiceRegistry(r *repository.RepoRegistry, storage storage.FileStorage) *ServiceRegistry {
	return &ServiceRegistry{
		Files:     files.NewFileService(r.Files, r.Shares, storage),
		Shares:    shares.NewShareService(r.Shares),
		Shortcuts: shortcuts.NewShortcutService(r.Files, r.Shares, r.Shortcuts),
	}
}

type FilesServiceMethods interface {
	UpdateFileName(ctx context.Context, in *inputs.UpdateFileNameInput) error
	GetSingleFileSummary(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) (*model.FileSummary, error)

	ListOwnedFiles(ctx context.Context, in *inputs.ListOwnedFilesInput) ([]*model.FileSummary, error)
	ListSharedFiles(ctx context.Context, in *inputs.ListSharedFilesInput) ([]*model.FileSummary, error)

	MakeFileCopy(ctx context.Context, in *inputs.MakeFileCopyInput) (*model.File, error)
	SoftDeleteFile(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) error

	// Download
	GetDownloadDetails(ctx context.Context, fildID uuid.UUID) (*model.DownloadRow, error)
}
type SharesServiceMethods interface {
	AddFileShares(ctx context.Context, in *inputs.AddFileSharesInput) error
	UpdateFileShare(ctx context.Context, in *inputs.UpdateFileShareInput) error
	RemoveFileShare(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID, recipientUserID uuid.UUID) error
	ListFileShares(ctx context.Context, fileID uuid.UUID, actorUserID uuid.UUID) ([]*model.FileShare, error)
	AddPublicAccess(ctx context.Context, in *inputs.AddPublicAccessInput) error
	RemovePublicAccess(ctx context.Context, in *inputs.RemovePublicAccessInput) error
}

type ShortcutsServiceMethods interface {
	CreateShortcut(ctx context.Context, in *inputs.CreateShortcutInput) (*model.FileShortcut, error)
	DeleteShortcut(ctx context.Context, in *inputs.DeleteShortcutInput) error
}
