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

type AddPartInput struct {
	UploadUUID uuid.UUID
	PartNumber int
	SizeBytes  int64
	Etag       string
}
