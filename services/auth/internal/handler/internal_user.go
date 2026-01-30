package handler

import (
	"encoding/json"
	"net/http"
)

// handler/internal_user.go
func (h *AuthHandler) ResolveUserIDsHandler(w http.ResponseWriter, r *http.Request) {
	var emails []string
	if err := json.NewDecoder(r.Body).Decode(&emails); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	result, err := h.service.ResolveUserIDs(r.Context(), emails)
	if err != nil {
		http.Error(w, "failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
