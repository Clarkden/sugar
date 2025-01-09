-- name: GetUser :one
SELECT
    *
FROM
    users
WHERE
    id = ?
LIMIT
    1;

-- name: ListUsers :many
SELECT
    *
FROM
    users;

-- name: CreateUser :one
INSERT INTO
    users (email, password)
VALUES
    (?, ?) RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET
    email = ?,
    password = ?
WHERE
    id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE
    id = ?;