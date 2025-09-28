-- name: CreateQuestion :one
INSERT INTO questions (id, classification_id, drug_id, text, choice_1, choice_2, choice_3, choice_4, explanation, answer_index)
VALUES (
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?,
    ?
)
RETURNING *;

-- name: DeleteQuestions :exec
DELETE FROM questions;

-- name: GetQuestionByID :one
SELECT * FROM questions
WHERE id = ?;

-- name: ListRandomQuestionsByClassification :many
SELECT * FROM questions
WHERE classification_id = ?
ORDER BY RANDOM()
LIMIT 5;