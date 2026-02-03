package postgres

import (
	"context"
	"database/sql"
	"errors"
	"upload/internal/model"
	"upload/shared"

	"github.com/lib/pq"
)

// type UploadChunkRepository interface {
// 	CreateChunk(chunk *model.UploadChunk) error
// 	GetSessionChunksCount(session_id int) (int, error)
// }

type PGChunkRepo struct {
	db *sql.DB
}

func NewChunkRepo(db *sql.DB) *PGChunkRepo {
	return &PGChunkRepo{db: db}
}

// queries
const (
	CreateChunkQuery           = `INSERT INTO upload_chunks (session_id, chunk_index, size_bytes, checksum) VALUES ($1, $2, $3, $4) RETURNING id`
	GetSessionChunksCountQuery = `Select Count(*) from upload_chunks where session_id = $1`
)

func (p *PGChunkRepo) CreateChunk(ctx context.Context, chunk *model.UploadChunk) error {
	err := p.db.QueryRowContext(
		ctx,
		CreateChunkQuery,
		chunk.SessionID,
		chunk.ChunkIndex,
		chunk.SizeBytes,
		chunk.CheckSum,
	).Scan(&chunk.ID)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			// duplicate chunk (idempotent retry)
			return shared.ErrChunkAlreadyExists
		}
		return err
	}

	return nil
}

func (p *PGChunkRepo) CountBySession(ctx context.Context, sessionID int64) (int, error) {
	var total int
	err := p.db.QueryRowContext(ctx, GetSessionChunksCountQuery, sessionID).Scan(&total)
	return total, err
}
