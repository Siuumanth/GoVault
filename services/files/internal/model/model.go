package model

import (
	"time"

	"github.com/google/uuid"
)

type FileShare struct {
	ID               int64
	FileID           int64
	SharedWithUserID uuid.UUID
	Permission       string
	CreatedAt        time.Time
}

type File struct {
	ID         int64
	FileUUID   uuid.UUID
	SessionID  *int64
	UserID     uuid.UUID
	FileName   string
	MimeType   *string
	SizeBytes  int64
	StorageKey string
	Checksum   *string
	CreatedAt  time.Time
	DeletedAt  *time.Time
}

type FileShortcut struct {
	ID        int64
	FileID    int64
	UserID    uuid.UUID
	CreatedAt time.Time
}

type PublicFile struct {
	FileID    int64
	CreatedAt time.Time
}
