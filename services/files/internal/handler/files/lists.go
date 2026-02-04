package files

import (
	"net/http"

	"files/internal/handler/common"
	"files/internal/handler/dto"
	"files/internal/service/inputs"
)

// GET /owned (private)
func (h *Handler) ListOwnedFiles(w http.ResponseWriter, r *http.Request) {
	actorID, err := common.GetRequiredActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	limit, offset := common.GetPagination(r)

	files, err := h.files.ListOwnedFiles(r.Context(), &inputs.ListOwnedFilesInput{
		UserID: *actorID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	resp := make([]dto.FileSummaryResponse, 0, len(files))
	for _, f := range files {
		resp = append(resp, dto.FileSummaryResponse{
			FileID:    f.FileUUID.String(),
			OwnerID:   f.UserID.String(),
			Name:      f.Name,
			MimeType:  f.MimeType,
			SizeBytes: f.SizeBytes,
			CreatedAt: f.CreatedAt,
		})
	}

	common.RespondJSON(w, http.StatusOK, dto.ListFilesResponse{Files: resp})
}

// GET /shared (Private)
func (h *Handler) ListSharedFiles(w http.ResponseWriter, r *http.Request) {
	actorID, err := common.GetRequiredActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	limit, offset := common.GetPagination(r)

	files, err := h.files.ListSharedFiles(r.Context(), &inputs.ListSharedFilesInput{
		UserID: *actorID,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	resp := make([]dto.FileSummaryResponse, 0, len(files))
	for _, f := range files {
		resp = append(resp, dto.FileSummaryResponse{
			FileID:    f.FileUUID.String(),
			OwnerID:   f.UserID.String(),
			Name:      f.Name,
			MimeType:  f.MimeType,
			SizeBytes: f.SizeBytes,
			CreatedAt: f.CreatedAt,
		})
	}

	common.RespondJSON(w, http.StatusOK, dto.ListFilesResponse{Files: resp})
}
