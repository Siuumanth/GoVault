package postgres

import (
	"context"
	model "files/internal/model"

	"github.com/google/uuid"
)

const FetchDownloadInfoQuery = `SELECT storage_key from files WHERE file_uuid = $1 AND deleted_at IS NULL`

func (r *FilesRepository) FetchDownloadInfo(ctx context.Context, fileID uuid.UUID) (*model.DownloadRow, error) {
	res := model.DownloadRow{}
	err := r.db.QueryRowContext(ctx, FetchDownloadInfoQuery, fileID).Scan(&res.StorageKey)
	return &res, err

}
