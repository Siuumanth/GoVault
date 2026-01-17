-- Database: govault

-- DROP DATABASE IF EXISTS govault;

CREATE DATABASE govault
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_India.1252'
    LC_CTYPE = 'English_India.1252'
    LOCALE_PROVIDER = 'libc'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;


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

CREATE TYPE upload_status AS ENUM (
    'pending', 
    'assembling', 
    'uploading', 
    'completed', 
    'failed'
);
-- ===============================
-- upload_sessions
-- ===============================
-- pending        // session created, no chunks yet
-- assembling     // all chunks received, merging
-- uploading      // file being sent to s3
-- completed      // file safely stored
-- failed         // unrecoverable error

-- ===============================
-- 2. upload_sessions
-- ===============================
CREATE TABLE upload_sessions (
    id BIGSERIAL PRIMARY KEY,
    upload_uuid UUID NOT NULL UNIQUE,
    user_id UUID NOT NULL REFERENCES users(id),

    file_name TEXT NOT NULL,
    file_size_bytes BIGINT NOT NULL,
    total_chunks INT NOT NULL,

    status upload_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);


-- ===============================
-- 3. upload_chunks
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
-- 4. files
-- ===============================
CREATE TABLE files (
    id BIGSERIAL PRIMARY KEY,
    file_uuid UUID NOT NULL UNIQUE,
    session_id BIGINT REFERENCES upload_sessions(id) ON DELETE SET NULL,
    user_id UUID NOT NULL REFERENCES users(id),

    file_name TEXT NOT NULL,
    mime_type TEXT,

    size_bytes BIGINT NOT NULL,
    storage_key TEXT NOT NULL,
    checksum TEXT,

    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

