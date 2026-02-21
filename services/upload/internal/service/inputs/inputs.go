package inputs

import (
	"io"

	"github.com/google/uuid"
)

type UploadSessionInput struct {
	UserID        uuid.UUID
	FileName      string
	FileSizeBytes int64
	UploadMethod  string
	PartSize      int64
}

type UploadChunkInput struct {
	UserID     uuid.UUID
	UploadUUID uuid.UUID
	ChunkID    int
	ChunkBytes io.Reader
	CheckSum   string
}

type S3UploadChunkInput struct {
	UserID     uuid.UUID
	UploadUUID uuid.UUID
	ChunkID    int
	CheckSum   string
}
