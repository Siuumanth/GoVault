package service

import (
	"upload/internal/model"
	"upload/internal/repository"

	"github.com/google/uuid"
)

type UploadService struct {
	registry repository.RepoRegistry
}

func NewService(registry repository.RepoRegistry) *UploadService {
	return &UploadService{registry: registry}
}

type ServiceMethods interface {
	UploadSession(inputs *UploadSessionInput) (*model.UploadSession, error)
	UploadChunk(inputs *UploadChunkInput) (*model.UploadChunk, error)
	GetUploadStatus(upload_uuid uuid.UUID) (*model.UploadSession, error)
}
