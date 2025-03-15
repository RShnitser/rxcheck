-- name: CreateAnswer :one
INSERT INTO answers (id, text, answer_order, question_id)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: DeleteAnswer :exec
DELETE FROM answers;

-- name: ListAnswersByQuestion :exec
SELECT * FROM answers
WHERE question_id = $1;