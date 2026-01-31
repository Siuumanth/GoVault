package dto

import "github.com/google/uuid"

type InternalFileMetadata struct {
	FileUUID   uuid.UUID `json:"file_uuid"`
	UserID     uuid.UUID `json:"user_id"`     // Owner
	StorageKey string    `json:"storage_key"` // S3 Path (Sensitive)
	MimeType   string    `json:"mime_type"`
	SizeBytes  int64     `json:"size_bytes"`
}
