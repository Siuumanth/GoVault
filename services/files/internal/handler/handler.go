package handler

import (
	"encoding/json"
	"errors"
	"files/internal/handler/dto"
	"files/internal/model"
	"files/internal/service"
	"files/internal/shared"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type Handler struct {
	registry *service.ServiceRegistry
}

func NewUploadHandler(regr *service.ServiceRegistry) *Handler {
	return &Handler{registry: regr}
}

// HELPERS:

func (h *Handler) getActorID(r *http.Request) (uuid.UUID, error) {
	uid, ok := r.Context().Value(shared.ActorIDKey).(uuid.UUID)
	if !ok {
		return uuid.Nil, shared.ErrUnauthorized
	}
	return uid, nil
}

func decodeJSON(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return err
	}

	// ensure only ONE JSON object
	if dec.More() {
		return errors.New("multiple JSON objects not allowed")
	}
	return nil
}

func respondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// --- Pagination Helper ---

func (h *Handler) getPagination(r *http.Request) (int, int) {
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = shared.PAGE_NO_DEFAULT
	}

	limit := shared.PAGE_LIMIT
	offset := (page - 1) * limit

	return limit, offset
}

func mapFileSummaries(files []*model.FileSummary) []dto.FileSummaryResponse {
	resp := make([]dto.FileSummaryResponse, 0, len(files))

	for _, f := range files {
		resp = append(resp, dto.FileSummaryResponse{
			FileID:    f.FileUUID.String(),
			OwnerID:   f.UserID.String(),
			Name:      f.Name,
			MimeType:  f.MimeType,
			Size:      f.SizeBytes,
			CreatedAt: f.CreatedAt,
		})
	}

	return resp
}
