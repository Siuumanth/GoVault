package repository

import (
	"upload/internal/model"

	"github.com/google/uuid"
)

type FileRepository interface {
	Create(file *model.File) error
	GetByID(fileID uuid.UUID) (*model.File, error)
}

// interface which stores methods for uploading chunks
type UploadChunkRepository interface {
	CreateChunk(chunk *model.UploadChunk) error
	GetTotalChunks(session_id int) int
}

type UploadSessionRepository interface {
	CreateUploadSession(session *model.UploadSession) error
	GetUploadSession(session_id int) (*model.UploadSession, error)
	GetTotalChunks(session_id int) int
	UpdateUploadStatus(session_id int, status string) error
}
