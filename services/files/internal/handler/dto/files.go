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
	SizeBytes int64     `json:"size_bytes"`
	CreatedAt time.Time `json:"created_at"`
}

type DownloadInfoResponse struct {
	DownloadURL string `json:"download_url"`
	ExpiresAt   int64  `json:"expires_at"`
}
