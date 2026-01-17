package service

import "github.com/google/uuid"

type UploadSessionInput struct {
	UserID        uuid.UUID
	FileName      string
	FileSizeBytes int64
}

type UploadChunkInput struct {
	UserID     uuid.UUID
	UploadUUID uuid.UUID
	ChunkID    int
	ChunkBytes []byte
	CheckSum   string
}
