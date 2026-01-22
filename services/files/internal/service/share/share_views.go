package share

import "github.com/google/uuid"

type FileShareParams struct {
	FileID           uuid.UUID
	SharedWithUserID uuid.UUID
	Permission       string
}

// TODO: make is shred with user ID
