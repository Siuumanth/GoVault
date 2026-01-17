package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"upload/internal/service"

	"github.com/google/uuid"
)

type Handler struct {
	uploadService *service.UploadService
}

func NewUploadHandler(uploadService *service.UploadService) *Handler {
	return &Handler{uploadService: uploadService}
}

// helpers
func userIDFromHeader(r *http.Request) (uuid.UUID, error) {
	id := r.Header.Get("X-User-ID")
	if id == "" {
		return uuid.Nil, fmt.Errorf("missing X-User-ID")
	}
	return uuid.Parse(id)
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
