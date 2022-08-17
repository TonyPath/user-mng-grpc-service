CREATE TABLE IF NOT EXISTS "users" (
    id                  UUID PRIMARY KEY,
    email               VARCHAR(255) NOT NULL UNIQUE,
    first_name          VARCHAR(255) NOT NULL,
    last_name           VARCHAR(255) NOT NULL,
    nickname            VARCHAR(255) NOT NULL,
    country             VARCHAR(2) NOT NULL,
    password            TEXT,
    created_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMPTZ
);
