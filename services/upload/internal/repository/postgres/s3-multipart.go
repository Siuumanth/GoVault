package postgres

import (
	"context"
	"database/sql"
	"upload/internal/model"
	"upload/shared"
)

// type PGMultipartRepo MultiPartMethods{
// 	GetPartsBySession(ctx context.Context, session_id int) ([]*model.UploadChunk, error)
//  CreatePart(ctx context.Context, part *model.UploadPart) error
// }

type PGMultipartRepo struct {
	db *sql.DB
}

func NewMultipartRepo(db *sql.DB) *PGMultipartRepo {
	return &PGMultipartRepo{db: db}
}

const CreatePartQuery = `
    INSERT INTO s3_multipart_parts (session_id, part_number, etag, size_bytes)
    VALUES ($1, $2, $3, $4)
    ON CONFLICT (session_id, part_number) DO NOTHING
`

func (p *PGMultipartRepo) CreatePart(ctx context.Context, part *model.UploadPart) error {
	res, err := p.db.ExecContext(
		ctx,
		CreatePartQuery,
		part.SessionID,
		part.PartNumber,
		part.Etag,
		part.SizeBytes,
	)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return shared.ErrPartAlreadyExists
	}

	return nil
}

const GetPartsBySessionQuery = `
    SELECT id, session_id, part_number, etag, size_bytes, uploaded_at
    FROM s3_multipart_parts
    WHERE session_id = $1
    ORDER BY part_number ASC
`

func (p *PGMultipartRepo) GetPartsBySession(ctx context.Context, sessionId int64) ([]*model.UploadPart, error) {
	rows, err := p.db.QueryContext(ctx, GetPartsBySessionQuery, sessionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parts []*model.UploadPart
	for rows.Next() {
		var part model.UploadPart
		err := rows.Scan(
			&part.ID,
			&part.SessionID,
			&part.PartNumber,
			&part.Etag,
			&part.SizeBytes,
			&part.UploadedAt,
		)
		if err != nil {
			return nil, err
		}
		parts = append(parts, &part)
	}

	return parts, rows.Err()
}
