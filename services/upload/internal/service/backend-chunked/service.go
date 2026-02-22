package backend

import (
	"upload/internal/clients"
	"upload/internal/repository"
	"upload/internal/storage"
)

type ProxyUploadService struct {
	registry   *repository.RepoRegistry
	storage    storage.FileStorage
	fileClient *clients.FileClient
}

func NewProxyUploadService(registry *repository.RepoRegistry, storage storage.FileStorage, fileClient *clients.FileClient) *ProxyUploadService {
	return &ProxyUploadService{
		registry:   registry,
		storage:    storage,
		fileClient: fileClient,
	}
}

// type ServiceMethods interface {
// 	UploadSession(ctx context.Context, inputs *UploadSessionInput) (*model.UploadSession, error)
// 	UploadChunk(ctx context.Context, inputs *UploadChunkInput) (*model.UploadChunk, error)
// 	GetUploadStatus(ctx context.Context, upload_uuid uuid.UUID) (*model.UploadSession, error)
// }
