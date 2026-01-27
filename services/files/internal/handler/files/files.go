package files

import (
	"errors"
	"net/http"

	"files/internal/handler/common"
	"files/internal/handler/dto"
	"files/internal/service"
	"files/internal/shared"

	"github.com/google/uuid"
)

// GET /{fileID}
func (h *Handler) GetSingleFileSummary(w http.ResponseWriter, r *http.Request) {
	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid file id", http.StatusBadRequest)
		return
	}

	actorID, _ := r.Context().Value(shared.ActorIDKey).(uuid.UUID)

	file, err := h.files.GetSingleFileSummary(r.Context(), fileID, actorID)
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

	common.RespondJSON(w, http.StatusOK, dto.FileSummaryResponse{
		FileID:    file.FileUUID.String(),
		OwnerID:   file.UserID.String(),
		Name:      file.Name,
		MimeType:  file.MimeType,
		Size:      file.SizeBytes,
		CreatedAt: file.CreatedAt,
	})
}

// PATCH /{fileID}
func (h *Handler) UpdateFileName(w http.ResponseWriter, r *http.Request) {
	actorID, err := common.GetActorID(r)
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
	if err := common.DecodeJSON(r, &req); err != nil || req.Name == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.files.UpdateFileName(r.Context(), &service.UpdateFileNameInput{
		FileID:      fileID,
		ActorUserID: actorID,
		NewName:     req.Name,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DELETE /{fileID}
func (h *Handler) SoftDeleteFile(w http.ResponseWriter, r *http.Request) {
	actorID, err := common.GetActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid file id", http.StatusBadRequest)
		return
	}

	if err := h.files.SoftDeleteFile(r.Context(), fileID, actorID); err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// POST /{fileID}/copy
func (h *Handler) MakeFileCopy(w http.ResponseWriter, r *http.Request) {
	actorID, err := common.GetActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid file id", http.StatusBadRequest)
		return
	}

	_, err = h.files.MakeFileCopy(r.Context(), &service.MakeFileCopyInput{
		FileID:      fileID,
		ActorUserID: actorID,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
