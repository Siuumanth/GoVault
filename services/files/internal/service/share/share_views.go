package share

import "github.com/google/uuid"

type ShareFileParams struct {
	FileID          uuid.UUID
	RecipientUserID uuid.UUID
	Permission      string
}
