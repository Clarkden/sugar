-- name: CreateSession :one
INSERT INTO sessions (user_id, session_id, created_at, expires_at) VALUES (?, ?, ?, ?) RETURNING *;

-- name: GetSessionByID :one
SELECT * FROM sessions WHERE session_id = ? LIMIT 1;

-- name: GetSessionsByuser_id :many
SELECT * FROM sessions WHERE user_id = ?;

-- name: DeleteSessionByID :exec
DELETE FROM sessions WHERE session_id = ?;

-- name: DeleteSessionsByuser_id :exec
DELETE FROM sessions WHERE user_id = ?;

-- name: UpdateSessionByID :one
UPDATE sessions SET user_id = ?, session_id = ?, created_at = ?, expires_at = ? WHERE user_id = ? RETURNING *;