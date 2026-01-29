package health

import "net/http"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Files service Working"))
}
