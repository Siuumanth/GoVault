
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE upload_status AS ENUM (
    'pending',
    'assembling',
    'uploading',
    'completed',
    'failed'
);

CREATE TABLE upload_sessions (
    id BIGSERIAL PRIMARY KEY,
    upload_uuid UUID NOT NULL UNIQUE,
    user_id UUID NOT NULL,

    file_name TEXT NOT NULL,
    file_size_bytes BIGINT NOT NULL,
    total_chunks INT NOT NULL,

    status upload_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_upload_sessions_uuid
ON upload_sessions(upload_uuid);

CREATE TABLE upload_chunks (
    id BIGSERIAL PRIMARY KEY,
    session_id BIGINT NOT NULL
        REFERENCES upload_sessions(id)
        ON DELETE CASCADE,

    chunk_index INT NOT NULL,
    size_bytes INT NOT NULL,
    checksum TEXT,

    uploaded_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (session_id, chunk_index)
);

CREATE INDEX idx_upload_chunks_session
ON upload_chunks(session_id);
