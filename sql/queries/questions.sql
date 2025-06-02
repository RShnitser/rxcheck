-- name: CreateQuestion :one
INSERT INTO questions (id, classification_id, drug_id, text, choice_1, choice_2, choice_3, choice_4, explanation, answer_index)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9
)
RETURNING *;

-- name: DeleteQuestions :exec
DELETE FROM questions;

-- name: ListRandomQuestionsByClassification :many
SELECT * FROM questions
WHERE classification_id = $1
ORDER BY RANDOM()
LIMIT 5;