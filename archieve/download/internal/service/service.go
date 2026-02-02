package service

import (
	"download/internal/repository"
	"download/internal/storage"
)

type DownloadService struct {
	repos   *repository.RepoRegistry
	storage storage.FileStorage
}

func NewDownloadService(
	repos *repository.RepoRegistry,
	storage storage.FileStorage,
) *DownloadService {
	return &DownloadService{
		repos:   repos,
		storage: storage,
	}
}
