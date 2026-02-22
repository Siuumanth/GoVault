package service

import (
	"upload/internal/clients"
	"upload/internal/repository"
	"upload/internal/service/backend-chunked"
	multipart "upload/internal/service/s3-multipart"
	"upload/internal/storage"
)

type ServiceRegistry struct {
	chunked   *backend.ProxyUploadService
	multipart *multipart.MultipartUploadService
}

func NewServiceRegistry(registry *repository.RepoRegistry, storage storage.FileStorage, fileClient *clients.FileClient) *ServiceRegistry {
	return &ServiceRegistry{
		chunked:   backend.NewProxyUploadService(registry, storage, fileClient),
		multipart: multipart.NewMultipartUploadService(registry, storage, fileClient),
	}
}
