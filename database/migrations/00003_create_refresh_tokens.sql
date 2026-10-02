-- +goose Up
CREATE TABLE refresh_tokens (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    hashed_token VARCHAR(500) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    expires_at TIMESTAMPTZ NOT NULL DEFAULT now() + INTERVAL '1 day',
    PRIMARY KEY (user_id, hashed_token)
);

-- +goose Down
DROP TABLE refresh_tokens;