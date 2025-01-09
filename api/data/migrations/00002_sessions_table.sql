-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    sessions (
        user_id INTEGER,
        session_id TEXT,
        created_at INTEGER,
        expires_at INTEGER,
        FOREIGN KEY (user_id) REFERENCES users (id)
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;

-- +goose StatementEnd