package repository

import (
	"database/sql"
	"upload/internal/repository/postgres"
)

type RepoRegistry struct {
	Sessions UploadSessionRepository
	Chunks   UploadChunkRepository
}

func NewRegistry(
	sessions UploadSessionRepository,
	chunks UploadChunkRepository,
) *RepoRegistry {
	return &RepoRegistry{
		Sessions: sessions,
		Chunks:   chunks,
	}
}

func NewRegistryFromDB(db *sql.DB) *RepoRegistry {
	return &RepoRegistry{
		Sessions: postgres.NewUploadSessionRepo(db),
		Chunks:   postgres.NewChunkRepo(db),
	}
}
