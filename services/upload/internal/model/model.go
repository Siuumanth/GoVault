package model

import (
	"time"

	"github.com/google/uuid"
)

// these are models returned by the repository and the actual schema

type UploadSession struct {
	ID          int       // internal BIGSERIAL
	UploadUUID  uuid.UUID // public ID
	UserID      uuid.UUID
	FileName    string
	FileSize    int64
	TotalChunks int
	Status      string
	CreatedAt   time.Time
}

type UploadChunk struct {
	ID         int
	SessionID  int
	ChunkIndex int
	SizeBytes  int
	CheckSum   string
	UploadedAt time.Time
}

type File struct {
	ID         int
	FileUUID   uuid.UUID
	UserID     uuid.UUID
	SessionID  int
	Name       string
	MimeType   string
	SizeBytes  int
	StorageKey string
	UploadedAt time.Time
}
