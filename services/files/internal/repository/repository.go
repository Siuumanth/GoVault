package repository

import (
	"files/internal/model"

	"github.com/google/uuid"
)

type FileRepository interface {
	CreateFile(file *model.File) error
	GetByID(fileID uuid.UUID) (*model.File, error)
}

// interface which stores methods for uploading chunks
type UploadChunkRepository interface {
	CreateChunk(chunk *model.UploadChunk) error
	CountBySession(session_id int) (int, error)
}

type UploadSessionRepository interface {
	CreateSession(session *model.UploadSession) error
	GetSessionByID(session_id int) (*model.UploadSession, error)
	GetSessionByUUID(upload_uuid uuid.UUID) (*model.UploadSession, error)
	UpdateSessionStatus(session_id int, status string) error
	DeleteSessionChunks(session_id int) error
}
