package dto

type ListFilesResponse struct {
	Files []FileSummaryResponse `json:"files"`
}
