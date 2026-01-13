package service

import (
	"upload/internal/model"
	"upload/internal/repository"

	"github.com/google/uuid"
)

type FileServiceMethods interface {
	GetFileByUUID(fileID uuid.UUID) (*model.File, error)
	GetFileByID(fileID uuid.UUID) (*model.File, error)
}

type FileService struct {
	registry *repository.Registry
}
