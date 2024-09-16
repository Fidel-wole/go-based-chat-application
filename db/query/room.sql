
-- name: CreateRoom :one
INSERT INTO rooms (name)
VALUES ($1)
RETURNING id, name;


