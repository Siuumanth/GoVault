package handler

import (
	"files/internal/handler/dto"
	"files/internal/service"
	"net/http"
)

// GET /moved
func (h *Handler) ListOwnedFiles(w http.ResponseWriter, r *http.Request) {
	actorID, err := h.getActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	limit, offset := h.getPagination(r)

	files, err := h.registry.Files.ListOwnedFiles(r.Context(), &service.ListOwnedFilesInput{
		UserID: actorID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	resp := dto.ListFilesResponse{
		Files: mapFileSummaries(files),
	}

	respondJSON(w, http.StatusOK, resp)
}

// GET /shared
func (h *Handler) ListSharedFiles(w http.ResponseWriter, r *http.Request) {
	actorID, err := h.getActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	limit, offset := h.getPagination(r)

	files, err := h.registry.Files.ListSharedFiles(r.Context(), &service.ListSharedFilesInput{
		UserID: actorID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	resp := dto.ListFilesResponse{
		Files: mapFileSummaries(files),
	}

	respondJSON(w, http.StatusOK, resp)
}
