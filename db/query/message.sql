-- name: CreateMessage :one
INSERT INTO messages (room_id, user_id, content)
VALUES ($1, $2, $3)
RETURNING id, room_id, user_id, content, created_at;

-- name: GetMessagesByRoom :many
SELECT id, room_id, user_id, content, created_at
FROM messages
WHERE room_id = $1
ORDER BY created_at ASC;