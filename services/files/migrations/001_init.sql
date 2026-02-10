
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











