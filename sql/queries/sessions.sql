-- name: CreateSession :one
INSERT INTO sessions (id, user_id, question_1, question_2, question_3, question_4, question_5, score, question_index)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    0,
    0
)
RETURNING *;

-- name: GetSessionByUserID :one
SELECT * FROM sessions
WHERE user_id = $1;

-- name: UpdateSession :exec
UPDATE sessions
SET score = $2, 
question_index = $3
WHERE id = $1;

-- name: DeleteSession :exec
DELETE FROM sessions WHERE user_id = $1;