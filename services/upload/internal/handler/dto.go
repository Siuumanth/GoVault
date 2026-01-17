package handler

import "github.com/google/uuid"

type UploadChunkRequest struct {
	UploadUUID uuid.UUID `json:"upload_uuid"`
	CheckSum   string    `json:"checksum"`
	ChunkBytes []byte    `json:"chunk_bytes"`
}

// stores dto to communicate from handler to service layer
type CreateUploadSessionRequest struct {
	UserID        uuid.UUID `json:"user_id"`
	FileName      string    `json:"file_name"`
	FileSizeBytes int64     `json:"file_size_bytes"`
}

// handler/dto/create_upload_session_response.go
type CreateUploadSessionResponse struct {
	UploadUUID  uuid.UUID `json:"upload_uuid"`
	TotalChunks int       `json:"total_chunks"`
}

type UploadStatusResponse struct {
	UploadUUID    string `json:"upload_uuid"`
	Status        string `json:"status"`
	TotalChunks   int    `json:"total_chunks"`
	UploadedCount int    `json:"uploaded_count"`
}
