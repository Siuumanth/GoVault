package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateFileRequest struct {
	FileUUID   uuid.UUID `json:"file_uuid"`
	UserID     uuid.UUID `json:"user_id"`
	Name       string    `json:"name"`
	SizeBytes  int64     `json:"size_bytes"`
	MimeType   string    `json:"mime_type"`
	CheckSum   string    `json:"checksum"`
	StorageKey string    `json:"storage_key"`
}

type UpdateFileNameRequest struct {
	Name string `json:"name"`
}

type FileSummaryResponse struct {
	FileID    string    `json:"file_id"`
	OwnerID   string    `json:"owner_id"`
	Name      string    `json:"name"`
	MimeType  string    `json:"mime_type"`
	SizeBytes int64     `json:"size_bytes"`
	CreatedAt time.Time `json:"created_at"`
}

type DownloadInfoResponse struct {
	DownloadURL string `json:"download_url"`
	ExpiresAt   int64  `json:"expires_at"`
}
