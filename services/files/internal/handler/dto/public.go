package dto

import "time"

type PublicAccessResponse struct {
	FileID string `json:"file_id"`
	Public bool   `json:"public"`
}
type CreateShortcutResponse struct {
	ShortcutID string    `json:"shortcut_id"`
	FileID     string    `json:"file_id"`
	CreatedAt  time.Time `json:"created_at"`
}
