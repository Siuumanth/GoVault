package service

import "upload/internal/repository"

type UploadService struct {
	registry repository.RepoRegistry
}

func NewService(registry repository.RepoRegistry) *UploadService {
	return &UploadService{registry: registry}
}

/*
Upload service methods:
*/
