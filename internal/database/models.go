// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Classification struct {
	ID   uuid.UUID
	Name string
}

type Drug struct {
	ID               uuid.UUID
	GenericName      string
	BrandName        string
	ClassificationID uuid.UUID
}

type Question struct {
	ID               uuid.UUID
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

type RefreshToken struct {
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	ExpiresAt time.Time
	RevokedAt sql.NullTime
}

type Session struct {
	ID            uuid.UUID
	UserID        uuid.UUID
	Question1     uuid.UUID
	Question2     uuid.UUID
	Question3     uuid.UUID
	Question4     uuid.UUID
	Question5     uuid.UUID
	Score         int32
	QuestionIndex int32
}

type User struct {
	ID             uuid.UUID
	UserName       string
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	LastDaily      sql.NullTime
	Streak         int32
}
