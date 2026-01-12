package repository

import "upload/internal/model"

type UploadFileRepository interface {
	Create(file *model.File) error
}

// interface which stores methods for uploading chunks
type UploadChunkRepository interface {
	CreateChunk(session_id int, chunk_id int, size_bytes int) error
	GetTotalChunks(session_id int) int
}

type UploadSessionRepository interface {
	CreateUploadSession(user_id string, file_name string, file_size int, total_chunks int) (int, error)
	GetUploadSession(session_id int) (int, error)
	UpdateUploadStatus(session_id int, status string) error
	SetUploadedChunks(session_id int, count int) error
}
