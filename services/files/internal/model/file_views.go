package model

import (
	"time"

	"github.com/google/uuid"
)

// internal/service/files_views.go
type FileSummary struct {
	FileUUID  uuid.UUID
	UserID    uuid.UUID
	Name      string
	MimeType  string
	SizeBytes int64
	CreatedAt time.Time
	IsPublic  bool
}

type CreateFileParams struct {
	UploadUUID *uuid.UUID
	FileUUID   uuid.UUID
	UserID     uuid.UUID
	Name       string
	MimeType   string
	SizeBytes  int64
	Checksum   *string
	StorageKey string
}

type DownloadRow struct {
	StorageKey string
	FileName   string
	MimeType   string
}
type DownloadResponse struct {
	DownloadURL string    `json:"download_url"`
	ExpiresAt   time.Time `json:"expires_at"`
	FileName    string    `json:"file_name"`
}
