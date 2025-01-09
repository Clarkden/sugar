-- +goose Up
CREATE TABLE
    users (
        id INTEGER PRIMARY KEY,
        email text NOT NULL,
        password text NOT NULL
    );

-- +goose Down
DROP TABLE users;