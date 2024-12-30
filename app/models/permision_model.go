package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe User object.
type Permission struct {
	ID          uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	Module      string    `db:"module" json:"module" validate:"required,lte=50"`
	Action      string    `db:"action" json:"action" validate:"required,lte=40"`
	Description string    `db:"description" json:"description" validate:"required,len=255"`
}
