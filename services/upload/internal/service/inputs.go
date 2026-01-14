package service

import "github.com/google/uuid"

type CreateUploadSessionInput struct {
	UserID        uuid.UUID
	FileName      string
	FileSizeBytes int64
}
