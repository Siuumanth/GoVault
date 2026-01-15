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
	GetSessionChunksCount(session_id int) (int, error)
}

type UploadSessionRepository interface {
	CreateSession(session *model.UploadSession) error
	GetSessionByID(session_id int) (*model.UploadSession, error)
	GetSessionByUUID(upload_uuid uuid.UUID) (int, error)
	UpdateSessionStatus(upload_uuid uuid.UUID, status string) error
	DeleteSessionChunks(session_id int) error
}
