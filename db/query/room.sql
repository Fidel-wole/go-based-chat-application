
-- name: CreateRoom :one
INSERT INTO rooms (name)
VALUES ($1)
RETURNING id, name;


-- name: GetRooms :many
SELECT id, name
FROM rooms;

-- name: GetRoom :one
SELECT id, name
FROM rooms
WHERE id = $1;
