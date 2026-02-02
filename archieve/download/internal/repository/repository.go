package repository

import (
	"download/internal/model"

	"github.com/google/uuid"
)

type FileRepository interface {
	GetByID(fileID uuid.UUID) (*model.File, error)
}
