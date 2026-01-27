package model

import "github.com/google/uuid"

type FileShareParams struct {
	FileID           uuid.UUID
	SharedWithUserID uuid.UUID
	Permission       string
}
