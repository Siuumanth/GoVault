package postgres

import (
	"database/sql"

	"download/internal/model"

	"github.com/google/uuid"
)

type PGFileRepo struct {
	db *sql.DB
}

func NewFileRepo(db *sql.DB) *PGFileRepo {
	return &PGFileRepo{db: db}
}

const GetFileQuery = `
SELECT file_uuid, user_id, file_name, mime_type, size_bytes, storage_key, created_at
FROM files
WHERE file_uuid = $1
`

func (p *PGFileRepo) GetByID(fileID uuid.UUID) (*model.File, error) {
	var file model.File
	err := p.db.QueryRow(GetFileQuery, fileID).Scan(
		&file.FileUUID,
		&file.UserID,
		&file.Name,
		&file.MimeType,
		&file.SizeBytes,
		&file.StorageKey,
		&file.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &file, nil
}
