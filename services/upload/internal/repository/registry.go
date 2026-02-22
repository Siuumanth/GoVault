package repository

import (
	"database/sql"
	"upload/internal/repository/postgres"
)

type RepoRegistry struct {
	Sessions UploadSessionRepository
	Chunks   UploadChunkRepository
	Parts    UploadPartRepository
}

func NewRegistry(
	sessions UploadSessionRepository,
	chunks UploadChunkRepository,
	// parts UploadPartRepository
	parts UploadPartRepository,
) *RepoRegistry {
	return &RepoRegistry{
		Sessions: sessions,
		Chunks:   chunks,
		Parts:    parts,
	}
}

func NewRegistryFromDB(db *sql.DB) *RepoRegistry {
	return &RepoRegistry{
		Sessions: postgres.NewUploadSessionRepo(db),
		Chunks:   postgres.NewChunkRepo(db),
		Parts:    postgres.NewMultipartRepo(db),
	}
}
