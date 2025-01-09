-- name: CreateSession :one
INSERT INTO sessions (userId, sessionId, createdAt, expiresAt) VALUES (?, ?, ?, ?) RETURNING *;

-- name: GetSessionByID :one
SELECT * FROM sessions WHERE sessionId = ? LIMIT 1;

-- name: GetSessionsByUserId :many
SELECT * FROM sessions WHERE userId = ?;

-- name: DeleteSessionByID :exec
DELETE FROM sessions WHERE sessionId = ?;

-- name: DeleteSessionsByUserID :exec
DELETE FROM sessions WHERE userId = ?;

-- name: UpdateSessionByID :one
UPDATE sessions SET userId = ?, sessionId = ?, createdAt = ?, expiresAt = ? WHERE userId = ? RETURNING *;