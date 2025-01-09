-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    sessions (
        userId INTEGER,
        sessionId TEXT,
        createdAt INTEGER,
        expiresAt INTEGER,
        FOREIGN KEY (userId) REFERENCES users (id)
    )
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE sessions;

-- +goose StatementEnd