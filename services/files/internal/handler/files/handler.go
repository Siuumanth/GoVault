package files

import "files/internal/service"

type Handler struct {
	files service.FilesService
}

func New(files service.FilesService) *Handler {
	return &Handler{files: files}
}
