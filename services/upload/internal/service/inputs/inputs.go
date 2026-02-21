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
}

type UploadChunkInput struct {
	UserID     uuid.UUID
	UploadUUID uuid.UUID
	ChunkID    int
	ChunkBytes io.Reader
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
