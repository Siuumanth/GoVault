
-- ============================================
-- AUTH DB
-- ============================================


CREATE DATABASE auth_db;
\c auth_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_users_email ON users(email);

-- ============================================
-- UPLOAD DB
-- ============================================

CREATE DATABASE upload_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Status tracking for the lifecycle of an upload
CREATE TYPE upload_status AS ENUM (
    'pending',
    'assembling',
    'uploading',
    'completed',
    'failed'
);

CREATE TYPE upload_method_enum AS ENUM (
    'proxy',
    'multipart'
);

CREATE TABLE upload_sessions (
    id BIGSERIAL PRIMARY KEY,
    upload_uuid UUID NOT NULL UNIQUE,
    user_id UUID NOT NULL, -- Foreign key to auth_db.users(id)

    file_name TEXT NOT NULL,
    file_size_bytes BIGINT NOT NULL,
    total_chunks INT NOT NULL,

    -- From Migration: Track which method was used and the external S3 ID
    upload_method upload_method_enum NOT NULL DEFAULT 'proxy',
    storage_upload_id TEXT, -- Used for S3 Multipart UploadId

    status upload_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_upload_sessions_uuid ON upload_sessions(upload_uuid);

-- ============================================
-- 1. PROXY METHOD TRACKING
-- ============================================

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

CREATE INDEX idx_upload_chunks_session ON upload_chunks(session_id);


-- ============================================
-- 2. S3 MULTIPART METHOD TRACKING
-- ============================================

CREATE TABLE s3_multipart_parts (
    id BIGSERIAL PRIMARY KEY,
    session_id BIGINT NOT NULL 
        REFERENCES upload_sessions(id) 
        ON DELETE CASCADE,

    part_number INT NOT NULL,
    etag TEXT NOT NULL, -- S3 specific identifier for completion
    size_bytes INT NOT NULL,

    uploaded_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (session_id, part_number)
);

CREATE INDEX idx_s3_parts_session ON s3_multipart_parts(session_id);


-- ============================================
-- FILES DB
-- ============================================

CREATE DATABASE files_db;
\c files_db;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE files (
    id BIGSERIAL PRIMARY KEY,
    file_uuid UUID NOT NULL UNIQUE,

    user_id UUID NOT NULL,

    file_name TEXT NOT NULL,
    mime_type TEXT,

    size_bytes BIGINT NOT NULL,
    storage_key TEXT NOT NULL,
    checksum TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP NULL
);
CREATE INDEX idx_active_files ON files (file_uuid) WHERE deleted_at IS NULL; 
-- for listing owned files 
CREATE INDEX idx_files_user_id
ON files(user_id);



CREATE TABLE file_shares (
    id BIGSERIAL PRIMARY KEY,
    file_uuid UUID NOT NULL
        REFERENCES files(file_uuid)
        ON DELETE CASCADE,

    shared_with_user_id UUID NOT NULL,
    permission TEXT NOT NULL CHECK (permission IN ('viewer', 'editor')),

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (file_uuid, shared_with_user_id)
);
-- composite index on file UUID and shared_with_user_id internally implemented 
-- for listing shared files
CREATE INDEX idx_file_shares_user
ON file_shares(shared_with_user_id);



CREATE TABLE file_shortcuts (
    id BIGSERIAL PRIMARY KEY,
    file_uuid UUID NOT NULL
        REFERENCES files(file_uuid)
        ON DELETE CASCADE,

    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    UNIQUE (file_uuid, user_id)
);
-- for fetching shortcuts
CREATE INDEX idx_file_shortcuts_users
ON file_shortcuts(user_id);

CREATE TABLE public_files (
    file_uuid UUID PRIMARY KEY
        REFERENCES files(file_uuid)
        ON DELETE CASCADE,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);










-- 1️⃣ Auth service DB
-- docker compose exec auth-service \
--   psql "$AUTH_POSTGRES_URL_DEV" -f migrations/001_init.sql

-- 2️⃣ Upload service DB
-- docker compose exec upload-service \
--   psql "$UPLOAD_POSTGRES_URL_DEV" -f migrations/001_init.sql

-- 3️⃣ Files service DB
-- docker compose exec files-service \
--   psql "$FILES_POSTGRES_URL_DEV" -f migrations/001_init.sql


