package service

import (
	"io"

	"github.com/google/uuid"
)

type UploadSessionInput struct {
	UserID        uuid.UUID
	FileName      string
	FileSizeBytes int64
}

type UploadChunkInput struct {
	UserID     uuid.UUID
	UploadUUID uuid.UUID
	ChunkID    int
	ChunkBytes io.Reader // ✅ stream
	CheckSum   string
}

// type CreateFileCommand struct {
// 	UploadUUID uuid.UUID
// 	UserID     uuid.UUID
// 	Name       string
// 	SizeBytes  int64
// 	MimeType   string
// 	CheckSum   string
// 	StorageKey string
// }
