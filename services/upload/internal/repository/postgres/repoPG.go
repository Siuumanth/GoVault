package postgres

import (
	"upload/internal/repository"
)

type Registry struct {
	Sessions repository.UploadSessionRepository
	Chunks   repository.UploadChunkRepository
	Files    repository.FileRepository
}

// func NewRegistry(db *sql.DB) *Registry {
// 	return &Registry{
// 		Sessions: NewUploadSessionRepo(db),
// 		Chunks:   NewUploadChunkRepo(db),
// 		Files:    NewFileRepo(db),
// 	}
// }
