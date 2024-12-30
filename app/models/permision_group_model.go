package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe User object.
type PermissionGroup struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	Name       string    `db:"name" json:"name" validate:"required,lte=150"`
	GroupType  string    `db:"group_type" json:"group_type" validate:"required,lte=50"`
	Permission string    `db:"user_status" json:"user_status" validate:"required,len=50"`
}
