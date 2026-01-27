package shortcuts

import "files/internal/service"

type Handler struct {
	shortcuts service.ShortcutsService
}

func New(shortcuts service.ShortcutsService) *Handler {
	return &Handler{shortcuts: shortcuts}
}
