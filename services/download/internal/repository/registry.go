package repository

import (
	"database/sql"

	"download/internal/repository/postgres"
)

type RepoRegistry struct {
	Files FileRepository
}

func NewRegistryFromDB(db *sql.DB) *RepoRegistry {
	return &RepoRegistry{
		Files: postgres.NewFileRepo(db),
	}
}
