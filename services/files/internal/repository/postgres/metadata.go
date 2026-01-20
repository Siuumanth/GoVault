package postgres

import (
	"context"
	"database/sql"
	"files/internal/service/files"

	"github.com/google/uuid"
)

/*
type MetaDataRepository interface {
	GetFileMetadata(fileID uuid.UUID) (*files.FileSummary, error)
	UpdateFileName(fileID uuid.UUID, newName string) error
}
*/

type MetaDataRepository struct {
	db *sql.DB
}

const (
	GetFileMetadataQuery = `
		SELECT
			file_uuid,
			user_id,
			name,
			mime_type,
			size,
			created_at
		FROM files
		WHERE file_uuid = $1
		AND deleted_at IS NULL

	`
)

func (r *MetaDataRepository) GetFileMetadata(
	ctx context.Context, fileUUID uuid.UUID,
) (*files.FileSummary, error) {

	var fs files.FileSummary

	err := r.db.QueryRowContext(
		ctx,
		GetFileMetadataQuery,
		fileUUID,
	).Scan(
		&fs.FileUUID,
		&fs.UserID,
		&fs.Name,
		&fs.MimeType,
		&fs.Size,
		&fs.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // service decides: not found vs no access
		}
		return nil, err
	}

	return &fs, nil
}

const UpdateFileNameQuery = `
		UPDATE files
		SET name = $2
		WHERE id = $1
		AND deleted_at IS NULL
	`

func (r *MetaDataRepository) UpdateFileName(
	ctx context.Context,
	fileUUID uuid.UUID,
	newName string,
) (bool, error) {

	result, err := r.db.ExecContext(
		ctx,
		UpdateFileNameQuery,
		fileUUID,
		newName,
	)
	if err != nil {
		return false, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	if rows == 0 {
		return false, nil // file not found or soft-deleted
	}

	return true, nil
}
