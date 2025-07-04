// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: sessions.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createSession = `-- name: CreateSession :one
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
RETURNING id, user_id, question_1, question_2, question_3, question_4, question_5, score, question_index
`

type CreateSessionParams struct {
	UserID    uuid.UUID
	Question1 uuid.UUID
	Question2 uuid.UUID
	Question3 uuid.UUID
	Question4 uuid.UUID
	Question5 uuid.UUID
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error) {
	row := q.db.QueryRowContext(ctx, createSession,
		arg.UserID,
		arg.Question1,
		arg.Question2,
		arg.Question3,
		arg.Question4,
		arg.Question5,
	)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Question1,
		&i.Question2,
		&i.Question3,
		&i.Question4,
		&i.Question5,
		&i.Score,
		&i.QuestionIndex,
	)
	return i, err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM sessions WHERE user_id = $1
`

func (q *Queries) DeleteSession(ctx context.Context, userID uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSession, userID)
	return err
}

const getSessionByUserID = `-- name: GetSessionByUserID :one
SELECT id, user_id, question_1, question_2, question_3, question_4, question_5, score, question_index FROM sessions
WHERE user_id = $1
`

func (q *Queries) GetSessionByUserID(ctx context.Context, userID uuid.UUID) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionByUserID, userID)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Question1,
		&i.Question2,
		&i.Question3,
		&i.Question4,
		&i.Question5,
		&i.Score,
		&i.QuestionIndex,
	)
	return i, err
}

const updateSession = `-- name: UpdateSession :exec
UPDATE sessions
SET score = $2, 
question_index = $3
WHERE id = $1
`

type UpdateSessionParams struct {
	ID            uuid.UUID
	Score         int32
	QuestionIndex int32
}

func (q *Queries) UpdateSession(ctx context.Context, arg UpdateSessionParams) error {
	_, err := q.db.ExecContext(ctx, updateSession, arg.ID, arg.Score, arg.QuestionIndex)
	return err
}
