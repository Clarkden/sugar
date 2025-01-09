-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    users (
        id INTEGER PRIMARY KEY,
        email text NOT NULL UNIQUE,
        password text NOT NULL
    );

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE users;

-- +goose StatementEnd