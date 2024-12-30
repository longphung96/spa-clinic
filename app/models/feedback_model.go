package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe User object.
type FeedBack struct {
	ID        uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	Email     string    `db:"email" json:"email" validate:"required,email,lte=255"`
}
