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
}

type CreateFileParams struct {
	SessionID  int
	FileUUID   uuid.UUID
	UserID     uuid.UUID
	Name       string
	MimeType   string
	SizeBytes  int64
	Checksum   string
	StorageKey string
}
