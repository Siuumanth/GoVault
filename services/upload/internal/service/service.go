package service

import (
	"upload/internal/clients"
	"upload/internal/repository"
	"upload/internal/storage"
)

type UploadService struct {
	registry   *repository.RepoRegistry
	storage    storage.FileStorage
	fileClient *clients.FileClient
}

func NewUploadService(registry *repository.RepoRegistry, storage storage.FileStorage, fileClient *clients.FileClient) *UploadService {
	return &UploadService{
		registry:   registry,
		storage:    storage,
		fileClient: fileClient,
	}
}

// type ServiceMethods interface {
// 	UploadSession(inputs *UploadSessionInput) (*model.UploadSession, error)
// 	UploadChunk(inputs *UploadChunkInput) (*model.UploadChunk, error)
// 	GetUploadStatus(upload_uuid uuid.UUID) (*model.UploadSession, error)
// }
