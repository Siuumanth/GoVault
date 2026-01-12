-- USERS
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- ===============================
-- ENUM TYPES
-- ===============================

CREATE TYPE upload_status AS ENUM ('pending', 'completed', 'failed');

-- ===============================
-- upload_sessions
-- ===============================
-- TODO: remove uploaded chunks and updated at 
CREATE TABLE upload_sessions (
    id BIGSERIAL PRIMARY KEY,

    upload_uuid UUID NOT NULL UNIQUE,
    user_id UUID NOT NULL References users(id),

    file_name TEXT NOT NULL,
    file_size_bytes BIGINT NOT NULL,
    total_chunks INT NOT NULL,
  --  uploaded_chunks INT NOT NULL DEFAULT 0,

    status upload_status NOT NULL DEFAULT 'pending',

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  -- updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- ===============================
-- upload_chunks
-- ===============================

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

-- ===============================
-- files
-- ===============================

CREATE TABLE files (
    id BIGSERIAL PRIMARY KEY,

    file_uuid UUID NOT NULL UNIQUE,
    user_id UUID NOT NULL References users(id),

    file_name TEXT NOT NULL,
    mime_type TEXT,

    size_bytes BIGINT NOT NULL,
    storage_key TEXT NOT NULL,
    checksum TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
