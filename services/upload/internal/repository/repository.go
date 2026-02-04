package repository

import (
	"context"
	"upload/internal/model"

	"github.com/google/uuid"
)

type FileRepository interface {
	//CreateFile(file *model.File) error
	//GetByID(fileID uuid.UUID) (*model.File, error)
}

// interface which stores methods for uploading chunks
type UploadChunkRepository interface {
	CreateChunk(ctx context.Context, chunk *model.UploadChunk) error
	CountBySession(ctx context.Context, session_id int64) (int, error)
}

type UploadSessionRepository interface {
	CreateSession(ctx context.Context, session *model.UploadSession) error
	GetSessionByID(ctx context.Context, session_id int64) (*model.UploadSession, error)
	GetSessionByUUID(ctx context.Context, upload_uuid uuid.UUID) (*model.UploadSession, error)
	UpdateSessionStatus(ctx context.Context, session_id int64, status string) error
	DeleteSessionChunks(ctx context.Context, session_id int64) error
}
