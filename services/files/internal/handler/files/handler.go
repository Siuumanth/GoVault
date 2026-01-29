package files

import (
	"files/internal/service/files"
)

type Handler struct {
	files *files.FileService
}

func NewHandler(files *files.FileService) *Handler {
	return &Handler{files: files}
}
