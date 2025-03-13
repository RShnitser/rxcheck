-- name: CreateQuestion :one
INSERT INTO questions (id, text, classification_id, drug_id)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3
)
RETURNING *;

-- name: DeleteQuestions :exec
DELETE FROM questions;

-- name: ListQuestionByClassification :many
SELECT * FROM questions
WHERE classification_id = $1;