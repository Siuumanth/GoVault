package shortcuts

import "files/internal/service/shortcuts"

type Handler struct {
	shortcuts *shortcuts.ShortcutService
}

func NewHandler(shortcuts *shortcuts.ShortcutService) *Handler {
	return &Handler{shortcuts: shortcuts}
}
