package repository

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
