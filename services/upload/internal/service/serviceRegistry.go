package service

import (
	"upload/internal/clients"
	"upload/internal/repository"
	"upload/internal/service/backend-chunked"
	multipart "upload/internal/service/s3-multipart"
	"upload/internal/storage"
)

type ServiceRegistry struct {
	Proxy     *backend.ProxyUploadService
	Multipart *multipart.MultipartUploadService
}

func NewServiceRegistry(registry *repository.RepoRegistry, storage storage.FileStorage, fileClient *clients.FileClient) *ServiceRegistry {
	return &ServiceRegistry{
		Proxy:     backend.NewProxyUploadService(registry, storage, fileClient),
		Multipart: multipart.NewMultipartUploadService(registry, storage, fileClient),
	}
}
