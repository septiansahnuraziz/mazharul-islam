-- +migrate Up notransaction
CREATE TYPE gender_enum AS ENUM ('male', 'female');

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    name varchar(155),
    age int,
    gender gender_enum DEFAULT NULL,
    location POINT DEFAULT NULL,
    interests text[],
    preferences JSONB DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX  "users_id_idx" ON "users" ("id");
CREATE INDEX  "users_preferences_idx" ON "users" ("preferences");
-- +migrate Down