package dto

// handler/dto/upload_status_response.go
type UploadStatusResponse struct {
	UploadUUID     string `json:"upload_uuid"`
	Status         string `json:"status"`
	TotalChunks    int    `json:"total_chunks"`
	UploadedChunks []int  `json:"uploaded_chunks"`
}
