package handler

import "github.com/google/uuid"

type UploadChunkRequest struct {
	UploadUUID uuid.UUID `json:"upload_uuid"`
	CheckSum   string    `json:"checksum"`
	ChunkBytes []byte    `json:"chunk_bytes"`
}

// stores dto to communicate from handler to service layer
type CreateUploadSessionRequest struct {
	FileName      string `json:"file_name"`
	FileSizeBytes int64  `json:"file_size_bytes"`
}

// handler/dto/create_upload_session_response.go
type CreateUploadSessionResponse struct {
	UploadUUID uuid.UUID `json:"upload_uuid"`
	TotalParts int64     `json:"total_chunks"`
}

type UploadStatusResponse struct {
	UploadUUID string `json:"upload_uuid"`
	Status     string `json:"status"`
	TotalParts int64  `json:"total_chunks"`
}

// --- Multi-Part Way ---
type CreateMultipartSessionRequest struct {
	FileName      string `json:"file_name"`
	FileSizeBytes int64  `json:"file_size_bytes"`
	PartSizeBytes int64  `json:"part_size_bytes"`
}

type AddS3PartRequest struct {
	UploadUUID uuid.UUID `json:"upload_uuid"`
	PartNumber int       `json:"part_number"`
	SizeBytes  int64     `json:"size_bytes"`
	Etag       string    `json:"etag"`
}

// POST /multipart/complete
type CompleteMultipartRequest struct {
	UploadUUID uuid.UUID `json:"upload_uuid"`
}
