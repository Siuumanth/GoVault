package clients

import "github.com/google/uuid"

type CreateFileRequest struct {
	FileUUID   uuid.UUID `json:"file_uuid"`
	UploadUUID uuid.UUID `json:"upload_uuid"`
	UserID     uuid.UUID `json:"user_id"`
	Name       string    `json:"name"`
	SizeBytes  int64     `json:"size_bytes"`
	MimeType   string    `json:"mime_type"`
	CheckSum   string    `json:"checksum"`
	StorageKey string    `json:"storage_key"`
}
