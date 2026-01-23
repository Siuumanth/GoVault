package postgres

import (
	"context"

	"github.com/google/uuid"
)

/*
type MetaDataRepository interface {
	GetFileMetadata(fileID uuid.UUID) (*model.FileSummary, error)
	UpdateFileName(fileID uuid.UUID, newName string) error
}
*/

const UpdateFileNameQuery = `
		UPDATE files
		SET name = $2
		WHERE id = $1
		AND deleted_at IS NULL
	`

func (r *FilesRepository) UpdateFileName(
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
