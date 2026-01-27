package postgres

import (
	"context"
	"database/sql"
	"errors"
	"files/internal/model"
	"files/internal/shared"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
)

const CreateShortcutQuery = `
INSERT INTO file_shortcuts (file_id, user_id)
VALUES ($1, $2)
RETURNING id, created_at
`

const DeleteShortcutQuery = `
DELETE FROM file_shortcuts
WHERE file_id = $1 AND user_id = $2
`

type ShortcutsRepository struct {
	db *sql.DB
}

func (r *ShortcutsRepository) CreateShortcut(
	ctx context.Context,
	fileID uuid.UUID,
	userID uuid.UUID,
) (*model.FileShortcut, error) {

	var sc model.FileShortcut

	err := r.db.QueryRowContext(
		ctx,
		CreateShortcutQuery,
		fileID,
		userID,
	).Scan(
		&sc.ID,
		&sc.CreatedAt,
	)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, shared.ErrRowExists
		}
		return nil, err
	}

	sc.FileID = fileID
	sc.UserID = userID

	return &sc, nil
}

func (r *ShortcutsRepository) DeleteShortcut(
	ctx context.Context,
	fileUUID uuid.UUID,
	userID uuid.UUID,
) error {
	_, err := r.db.ExecContext(
		ctx,
		DeleteShortcutQuery,
		fileUUID,
		userID,
	)
	return err
}
