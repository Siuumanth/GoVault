package postgres

import (
	"context"
	model "files/internal/model"

	"github.com/google/uuid"
)

const ListOwnedFilesQuery = `
SELECT f.file_uuid, f.user_id, f.file_name, f.mime_type, f.size_bytes, f.created_at,
       (p.file_uuid IS NOT NULL) as is_public
FROM files f
LEFT JOIN public_files p ON f.file_uuid = p.file_uuid
WHERE f.user_id = $1 AND f.deleted_at IS NULL
LIMIT $2 OFFSET $3
`

func (r *FilesRepository) FetchOwnedFiles(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*model.FileSummary, error) {
	var fs []*model.FileSummary

	rows, err := r.db.QueryContext(ctx, ListOwnedFilesQuery, userID, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		f := new(model.FileSummary)
		if err := rows.Scan(
			&f.FileUUID,
			&f.UserID,
			&f.Name,
			&f.MimeType,
			&f.SizeBytes,
			&f.CreatedAt,
			&f.IsPublic,
		); err != nil {
			return nil, err
		}
		fs = append(fs, f)
	}
	return fs, nil
}

const ListSharedFilesQuery = `
SELECT 
    f.file_uuid, 
    f.user_id, 
    f.file_name, 
    f.mime_type, 
    f.size_bytes, 
    f.created_at,
    (p.file_uuid IS NOT NULL) as is_public
FROM files f
JOIN file_shares FS ON FS.file_uuid = f.file_uuid
LEFT JOIN public_files p ON f.file_uuid = p.file_uuid
WHERE FS.user_id = $1
AND f.deleted_at IS NULL
LIMIT $2 OFFSET $3
`

func (r *FilesRepository) FetchSharedFiles(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*model.FileSummary, error) {
	var fs []*model.FileSummary

	rows, err := r.db.QueryContext(ctx, ListSharedFilesQuery, userID, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		f := new(model.FileSummary)
		if err := rows.Scan(
			&f.FileUUID,
			&f.UserID,
			&f.Name,
			&f.MimeType,
			&f.SizeBytes,
			&f.CreatedAt,
			&f.IsPublic, // Added scan for public status
		); err != nil {
			return nil, err
		}
		fs = append(fs, f)
	}
	return fs, nil
}
