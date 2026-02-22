ALTER TABLE upload_sessions
ADD COLUMN upload_method TEXT NOT NULL DEFAULT 'proxy',
ADD COLUMN storage_upload_id TEXT;

-- Session for multipart upload:
-- pending (no parts gathered yet)  -> uploading -> completed/failed

CREATE TABLE s3_multipart_parts (
    id BIGSERIAL PRIMARY KEY,
    session_id BIGINT NOT NULL
        REFERENCES upload_sessions(id)
        ON DELETE CASCADE,

    part_number INT NOT NULL,
    etag TEXT NOT NULL,
    size_bytes INT NOT NULL,

    uploaded_at TIMESTAMP NOT NULL DEFAULT NOW(),

    UNIQUE (session_id, part_number)
);

CREATE INDEX idx_s3_parts_session
ON s3_multipart_parts(session_id);