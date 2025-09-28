-- name: CreateSession :one
INSERT INTO sessions (id, user_id, question_1, question_2, question_3, question_4, question_5, score, question_index)
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    0,
    0
)
RETURNING *;

-- name: GetSessionByUserID :one
SELECT * FROM sessions
WHERE user_id = ?;

-- name: UpdateSession :exec
UPDATE sessions
SET score = ?, 
question_index = ?
WHERE id = ?;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE user_id = ?;