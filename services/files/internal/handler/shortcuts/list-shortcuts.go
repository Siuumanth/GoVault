package shortcuts

import (
	"files/internal/handler/common"
	"files/internal/handler/dto"
	"files/internal/service/inputs"
	"net/http"
)

// GET /shortcuts (Private)
func (h *Handler) ListShortcuts(w http.ResponseWriter, r *http.Request) {
	actorID, err := common.GetRequiredActorID(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	limit, offset := common.GetPagination(r)

	files, err := h.shortcuts.ListUsersShortcutsWithFiles(r.Context(), &inputs.ListUsersShortcutsWithFilesInput{
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
