-- name: CreateQuiz :one
INSERT INTO quizzes (id, user_id, question_1, question_2, question_3, question_4, question_5, score)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    0
)
RETURNING *;

-- name: GetQuizByUserID :one
SELECT * FROM quizzes
WHERE user_id = $1;

-- name: UpdateQuizScore :exec
UPDATE quizzes SET score = $2
WHERE id = $1;

-- name: DeleteQuiz :exec
DELETE FROM quizzes WHERE user_id = $1;