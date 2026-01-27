package handler

import (
	"errors"
	"files/internal/handler/dto"
	"files/internal/service"
	"files/internal/shared"
	"net/http"

	"github.com/google/uuid"
)

// POST /{fileID}/shares
func (h *Handler) AddFileShares(w http.ResponseWriter, r *http.Request) {
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

	var req dto.AddFileSharesRequest
	if err := decodeJSON(r, &req); err != nil || len(req.Recipients) == 0 {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	recipients := make([]service.ShareRecipientInput, 0, len(req.Recipients))
	for _, r := range req.Recipients {
		if r.Email == "" || r.Permission == "" {
			http.Error(w, "invalid recipient data", http.StatusBadRequest)
			return
		}
		recipients = append(recipients, service.ShareRecipientInput{
			Email:      r.Email,
			Permission: r.Permission,
		})
	}

	err = h.registry.Sharing.AddFileShares(r.Context(), &service.AddFileSharesInput{
		FileID:      fileID,
		ActorUserID: actorID,
		Recipients:  recipients,
	})
	if err != nil {
		switch {
		case errors.Is(err, shared.ErrTooManyShares):
			http.Error(w, err.Error(), http.StatusBadRequest)
		case errors.Is(err, shared.ErrUnauthorized):
			http.Error(w, err.Error(), http.StatusForbidden)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// PATCH /{fileID}/shares/{userID}
func (h *Handler) UpdateFileShare(w http.ResponseWriter, r *http.Request) {
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

	userID, err := uuid.Parse(r.PathValue("userID"))
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	var req dto.UpdateFileShareRequest
	if err := decodeJSON(r, &req); err != nil || req.Permission == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.registry.Sharing.UpdateFileShare(r.Context(), &service.UpdateFileShareInput{
		FileID:          fileID,
		ActorUserID:     actorID,
		RecipientUserID: userID,
		Permission:      req.Permission,
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

// GET /{fileID}/shares
func (h *Handler) ListFileShares(w http.ResponseWriter, r *http.Request) {
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

	shares, err := h.registry.Sharing.ListFileShares(r.Context(), fileID, actorID)
	if err != nil {
		switch {
		case errors.Is(err, shared.ErrUnauthorized):
			http.Error(w, err.Error(), http.StatusForbidden)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	resp := make([]dto.FileShareResponse, 0, len(shares))
	for _, s := range shares {
		resp = append(resp, dto.FileShareResponse{
			UserID:     s.SharedWithUserID.String(),
			Permission: s.Permission,
			CreatedAt:  s.CreatedAt,
		})
	}

	respondJSON(w, http.StatusOK, resp)
}
