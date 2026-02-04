package common

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"files/internal/shared"

	"github.com/google/uuid"
)

func GetOptionalActorID(r *http.Request) *uuid.UUID {
	if uid, ok := r.Context().Value(shared.ActorIDKey).(uuid.UUID); ok {
		return &uid
	}
	return nil
}

func GetRequiredActorID(r *http.Request) (*uuid.UUID, error) {
	uid, ok := r.Context().Value(shared.ActorIDKey).(uuid.UUID)
	if !ok {
		return nil, shared.ErrUnauthorized
	}
	return &uid, nil
}

func DecodeJSON(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return err
	}
	if dec.More() {
		return errors.New("multiple json objects")
	}
	return nil
}

func RespondJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func GetPagination(r *http.Request) (limit, offset int) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = shared.PAGE_NO_DEFAULT
	}

	limit = shared.PAGE_LIMIT
	offset = (page - 1) * limit
	return
}
