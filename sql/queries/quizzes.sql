-- name: CreateQuiz :one
INSERT INTO quizzes (id, user_id, question_1, question_2, question_3, question_4, question_5, next_question_index)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    &6,
    0
)
RETURNING *;

-- name: DeleteQuiz :exec
DELETE FROM quizzes WHERE user_id = $1;