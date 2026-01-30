package model

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID         int
	FileUUID   uuid.UUID
	UserID     uuid.UUID
	Name       string
	MimeType   string
	SizeBytes  int64
	StorageKey string
	CreatedAt  time.Time
}
