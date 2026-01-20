package files

import (
	"time"

	"github.com/google/uuid"
)

// internal/service/files_views.go
type FileSummary struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Name      string
	MimeType  string
	Size      int64
	CreatedAt time.Time
}

type CreateFileParams struct {
	ID         uuid.UUID
	UserID     uuid.UUID
	Name       string
	MimeType   string
	Size       int64
	Checksum   string
	StorageKey string
}
