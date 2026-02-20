package shares

import (
	"errors"
	"net/http"

	"files/internal/handler/common"
	"files/internal/handler/dto"
	"files/internal/service/inputs"
	"files/internal/shared"

	"github.com/google/uuid"
)

// POST /{fileID}/shares (private)
func (h *Handler) AddFileShares(w http.ResponseWriter, r *http.Request) {
	actorID, err := common.GetRequiredActorID(r)
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
	if err := common.DecodeJSON(r, &req); err != nil || len(req.Recipients) == 0 {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	recipients := make([]inputs.ShareRecipientInput, 0, len(req.Recipients))
	for _, r := range req.Recipients {
		if r.Email == "" || r.Permission == "" {
			http.Error(w, "invalid recipient", http.StatusBadRequest)
			return
		}
		recipients = append(recipients, inputs.ShareRecipientInput{
			Email:      r.Email,
			Permission: r.Permission,
		})
	}

	err = h.shares.AddFileShares(r.Context(), &inputs.AddFileSharesInput{
		FileID:      fileID,
		ActorUserID: *actorID,
		Recipients:  recipients,
	})
	if err != nil {
		switch {
		case errors.Is(err, shared.ErrTooManyShares):
			http.Error(w, err.Error(), http.StatusBadRequest)
		default:
			http.Error(w, "internal error : "+err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// PATCH /{fileID}/shares/{userID} (private)
func (h *Handler) UpdateFileShare(w http.ResponseWriter, r *http.Request) {
	actorID, err := common.GetRequiredActorID(r)
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
	if err := common.DecodeJSON(r, &req); err != nil || req.Permission == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err = h.shares.UpdateFileShare(r.Context(), &inputs.UpdateFileShareInput{
		FileID:          fileID,
		ActorUserID:     *actorID,
		RecipientUserID: userID,
		Permission:      req.Permission,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DELETE /{fileID}/shares/{userID} (private)
func (h *Handler) RemoveFileShare(w http.ResponseWriter, r *http.Request) {
	actorID, err := common.GetRequiredActorID(r)
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

	err = h.shares.RemoveFileShare(r.Context(), fileID, *actorID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /{fileID}/shares (private)
func (h *Handler) ListFileShares(w http.ResponseWriter, r *http.Request) {
	actorID, err := common.GetRequiredActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	fileID, err := uuid.Parse(r.PathValue("fileID"))
	if err != nil {
		http.Error(w, "invalid file id", http.StatusBadRequest)
		return
	}

	shares, err := h.shares.ListFileShares(r.Context(), fileID, *actorID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
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

	common.RespondJSON(w, http.StatusOK, resp)
}
