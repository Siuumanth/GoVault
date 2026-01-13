package postgres

import (
	"database/sql"
	"upload/internal/model"
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
	CreateChunkQuery           = `INSERT INTO chunks (session_id, chunk_index, size_bytes, checksum) VALUES ($1, $2, $3, $4) RETURNING id`
	GetSessionChunksCountQuery = `Select Count(*) from chunks where session_id = $1`
)

func (p *PGChunkRepo) CreateChunk(chunk *model.UploadChunk) error {
	// we have CreateChunkQuerys
	err := p.db.QueryRow(
		CreateChunkQuery,
		chunk.SessionID,
		chunk.ChunkIndex,
		chunk.SizeBytes,
		chunk.CheckSum,
	).Scan(&chunk.ID)

	return err
}

func (p *PGChunkRepo) CountBySession(sessionID int) (int, error) {
	var total int
	err := p.db.QueryRow(GetSessionChunksCountQuery, sessionID).Scan(&total)
	return total, err
}
