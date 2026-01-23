package repository

type RepoRegistry struct {
	File      FileRepository
	Sharing   ShareRepository
	Shortcuts ShortcutsRepository
}

func NewRegistry(
	files FileRepository,
	sharing ShareRepository,
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
