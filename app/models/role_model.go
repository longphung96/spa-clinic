package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe User object.
type Role struct {
	ID               uuid.UUID       `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt        time.Time       `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time       `db:"updated_at" json:"updated_at"`
	Name             string          `db:"name" json:"name" validate:"required,lte=50"`
	Tenant           string          `db:"tenant" json:"tenant," validate:"required,lte=255"`
	Permissions      []Permission    `db:"permission" json:"permission" validate:"required,len=150"`
	PermissionGroups PermissionGroup `db:"permission_group" json:"permission_group" validate:"required,lte=25"`
}
