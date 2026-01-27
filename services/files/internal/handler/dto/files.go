package dto

import "time"

type UpdateFileNameRequest struct {
	Name string `json:"name"`
}

type FileSummaryResponse struct {
	FileID    string    `json:"file_id"`
	OwnerID   string    `json:"owner_id"`
	Name      string    `json:"name"`
	MimeType  string    `json:"mime_type"`
	Size      int64     `json:"size"`
	CreatedAt time.Time `json:"created_at"`
}
