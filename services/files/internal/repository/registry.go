package repository

import (
	"database/sql"
	"files/internal/repository/postgres"
)

type RepoRegistry struct {
	Files     FilesRepository
	Shares    SharesRepository
	Shortcuts ShortcutsRepository
}

func NewRegistry(
	files FilesRepository,
	shares SharesRepository,
	shortcuts ShortcutsRepository,
) *RepoRegistry {
	return &RepoRegistry{
		Files:     files,
		Shares:    shares,
		Shortcuts: shortcuts,
	}
}

func NewPostgresRegistry(db *sql.DB) *RepoRegistry {
	return &RepoRegistry{
		Files:     postgres.NewFilesRepository(db),
		Shares:    postgres.NewFileShareRepository(db),
		Shortcuts: postgres.NewShortcutsRepository(db),
	}
}
