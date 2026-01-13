package repository

type Registry struct {
	Sessions UploadSessionRepository
	Chunks   UploadChunkRepository
	Files    FileRepository
}

func NewRegistry(
	sessions UploadSessionRepository,
	chunks UploadChunkRepository,
	files FileRepository,
) *Registry {
	return &Registry{
		Sessions: sessions,
		Chunks:   chunks,
		Files:    files,
	}
}
