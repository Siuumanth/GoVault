package repository

import (
	"database/sql"
	"upload/internal/repository/postgres"
)

type RepoRegistry struct {
	Sessions UploadSessionRepository
	Chunks   UploadChunkRepository
	Files    FileRepository
}

func NewRegistry(
	sessions UploadSessionRepository,
	chunks UploadChunkRepository,
	files FileRepository,
) *RepoRegistry {
	return &RepoRegistry{
		Sessions: sessions,
		Chunks:   chunks,
		Files:    files,
	}
}

func NewRegistryFromDB(db *sql.DB) *RepoRegistry {
	return &RepoRegistry{
		Sessions: postgres.NewUploadSessionRepo(db),
		Chunks:   postgres.NewChunkRepo(db),
		Files:    postgres.NewFileRepo(db),
	}
}
