package dto

type UploadChunkMeta struct {
	UploadUUID string
	Index      int
	Checksum   string
}

type UploadChunkResponse struct {
	Status string `json:"status"`
}
