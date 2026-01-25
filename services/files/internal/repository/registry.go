package repository

type RepoRegistry struct {
	File      FilesRepository
	Sharing   SharesRepository
	Shortcuts ShortcutsRepository
}

func NewRegistry(
	files FilesRepository,
	sharing SharesRepository,
	shortcuts ShortcutsRepository,
) *RepoRegistry {
	return &RepoRegistry{
		File:      files,
		Sharing:   sharing,
		Shortcuts: shortcuts,
	}
}

// func NewRegistryFromDB(db *sql.DB) *RepoRegistry {
// 	return &RepoRegistry{
// 		Metadata:  postgres.NewMetaDataRepo(db),
// 		File:      postgres.NewFileRepo(db),
// 		Sharing:   postgres.NewShareRepo(db),
// 		Shortcuts: postgres.NewShortcutsRepo(db),
// 	}
// }
