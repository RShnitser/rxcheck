// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: questions.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createQuestion = `-- name: CreateQuestion :one
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
RETURNING id, classification_id, drug_id, text, choice_1, choice_2, choice_3, choice_4, explanation, answer_index
`

type CreateQuestionParams struct {
	ClassificationID uuid.UUID
	DrugID           uuid.UUID
	Text             string
	Choice1          string
	Choice2          string
	Choice3          string
	Choice4          string
	Explanation      string
	AnswerIndex      int32
}

func (q *Queries) CreateQuestion(ctx context.Context, arg CreateQuestionParams) (Question, error) {
	row := q.db.QueryRowContext(ctx, createQuestion,
		arg.ClassificationID,
		arg.DrugID,
		arg.Text,
		arg.Choice1,
		arg.Choice2,
		arg.Choice3,
		arg.Choice4,
		arg.Explanation,
		arg.AnswerIndex,
	)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.ClassificationID,
		&i.DrugID,
		&i.Text,
		&i.Choice1,
		&i.Choice2,
		&i.Choice3,
		&i.Choice4,
		&i.Explanation,
		&i.AnswerIndex,
	)
	return i, err
}

const deleteQuestions = `-- name: DeleteQuestions :exec
DELETE FROM questions
`

func (q *Queries) DeleteQuestions(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteQuestions)
	return err
}

const getQuestionByID = `-- name: GetQuestionByID :one
SELECT id, classification_id, drug_id, text, choice_1, choice_2, choice_3, choice_4, explanation, answer_index FROM questions
WHERE id = $1
`

func (q *Queries) GetQuestionByID(ctx context.Context, id uuid.UUID) (Question, error) {
	row := q.db.QueryRowContext(ctx, getQuestionByID, id)
	var i Question
	err := row.Scan(
		&i.ID,
		&i.ClassificationID,
		&i.DrugID,
		&i.Text,
		&i.Choice1,
		&i.Choice2,
		&i.Choice3,
		&i.Choice4,
		&i.Explanation,
		&i.AnswerIndex,
	)
	return i, err
}

const listRandomQuestionsByClassification = `-- name: ListRandomQuestionsByClassification :many
SELECT id, classification_id, drug_id, text, choice_1, choice_2, choice_3, choice_4, explanation, answer_index FROM questions
WHERE classification_id = $1
ORDER BY RANDOM()
LIMIT 5
`

func (q *Queries) ListRandomQuestionsByClassification(ctx context.Context, classificationID uuid.UUID) ([]Question, error) {
	rows, err := q.db.QueryContext(ctx, listRandomQuestionsByClassification, classificationID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Question
	for rows.Next() {
		var i Question
		if err := rows.Scan(
			&i.ID,
			&i.ClassificationID,
			&i.DrugID,
			&i.Text,
			&i.Choice1,
			&i.Choice2,
			&i.Choice3,
			&i.Choice4,
			&i.Explanation,
			&i.AnswerIndex,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
