package handler

import (
	"errors"
	"files/internal/handler/dto"
	"files/internal/service"
	"files/internal/shared"
	"net/http"

	"github.com/google/uuid"
)

// GET /{fileID}
func (h *Handler) GetSingleFileSummary(w http.ResponseWriter, r *http.Request) {
	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid file id", http.StatusBadRequest)
		return
	}

	// actor is optional (public access)
	actorID, _ := r.Context().Value(shared.ActorIDKey).(uuid.UUID)

	file, err := h.registry.Files.GetSingleFileSummary(r.Context(), fileID, actorID)
	if err != nil {
		switch {
		case errors.Is(err, shared.ErrRowNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
		case errors.Is(err, shared.ErrUnauthorized):
			http.Error(w, err.Error(), http.StatusForbidden)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	resp := dto.FileSummaryResponse{
		FileID:    file.FileUUID.String(),
		OwnerID:   file.UserID.String(),
		Name:      file.Name,
		MimeType:  file.MimeType,
		Size:      file.SizeBytes,
		CreatedAt: file.CreatedAt,
	}

	respondJSON(w, http.StatusOK, resp)
}

// PATCH /{fileID}
func (h *Handler) UpdateFileName(w http.ResponseWriter, r *http.Request) {
	actorID, err := h.getActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid file id", http.StatusBadRequest)
		return
	}

	var req dto.UpdateFileNameRequest
	if err := decodeJSON(r, &req); err != nil || req.Name == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.registry.Files.UpdateFileName(r.Context(), &service.UpdateFileNameInput{
		FileID:      fileID,
		ActorUserID: actorID,
		NewName:     req.Name,
	})
	if err != nil {
		switch {
		case errors.Is(err, shared.ErrUnauthorized):
			http.Error(w, err.Error(), http.StatusForbidden)
		case errors.Is(err, shared.ErrRowNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DELETE /{fileID}
func (h *Handler) SoftDeleteFile(w http.ResponseWriter, r *http.Request) {
	actorID, err := h.getActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid file id", http.StatusBadRequest)
		return
	}

	err = h.registry.Files.SoftDeleteFile(r.Context(), fileID, actorID)
	if err != nil {
		switch {
		case errors.Is(err, shared.ErrUnauthorized):
			http.Error(w, err.Error(), http.StatusForbidden)
		case errors.Is(err, shared.ErrRowNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// POST /{fileID}/copy
func (h *Handler) MakeFileCopy(w http.ResponseWriter, r *http.Request) {
	actorID, err := h.getActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid file id", http.StatusBadRequest)
		return
	}

	_, err = h.registry.Files.MakeFileCopy(r.Context(), &service.MakeFileCopyInput{
		FileID:      fileID,
		ActorUserID: actorID,
	})
	if err != nil {
		switch {
		case errors.Is(err, shared.ErrUnauthorized):
			http.Error(w, err.Error(), http.StatusForbidden)
		case errors.Is(err, shared.ErrRowNotFound):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}
