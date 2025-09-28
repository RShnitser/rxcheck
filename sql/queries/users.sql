-- name: CreateUser :one
INSERT INTO users (id, user_name, hashed_password, created_at, updated_at, last_daily, streak)
VALUES (
    ?,
    ?,
    ?,
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP,
    NULL,
    0
)
RETURNING *;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUserByUserName :one
SELECT * FROM users
WHERE user_name = ?;