// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, email, hashed_password, created_at, updated_at, last_daily, streak)
VALUES (
    gen_random_uuid(),
    $1,
    $2,
    NOW(),
    NOW(),
    NULL,
    0
)
RETURNING id, email, hashed_password, created_at, updated_at, last_daily, streak
`

type CreateUserParams struct {
	Email          string
	HashedPassword string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Email, arg.HashedPassword)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastDaily,
		&i.Streak,
	)
	return i, err
}
