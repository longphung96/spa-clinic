package models

import (
	"time"

	"github.com/google/uuid"
)

// User struct to describe User object.
type Customer struct {
	ID                  uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt           time.Time `db:"created_at" json:"created_at"`
	UpdatedAt           time.Time `db:"updated_at" json:"updated_at"`
	Email               string    `db:"email" json:"email" validate:"required,email,lte=255"`
	Phone               string    `db:"phone" json:"phone" validate:"required,lte=11"`
	Title               string    `db:"title" json:"title" validate:"required,len=10"`
	Gender              string    `db:"gender" json:"gender" validate:"required,lte=10"`
	Code                string    `db:"code" json:"code" validate:"required,lte=20"`
	Name                string    `db:"name" json:"name" validate:"required,lte=50"`
	AsciiName           string    `db:"ascii_name" json:"ascii_name" validate:"required,lte=50"`
	Address             string    `db:"address" json:"address" validate:"required,lte=100"`
	District            string    `db:"district" json:"district" validate:"required,lte=100"`
	City                string    `db:"city" json:"city" validate:"required,lte=20"`
	DayOfBirth          int16     `db:"day_of_birth" json:"day_of_birth" validate:"required,lte=50"`
	MonthOfBirth        int16     `db:"month_of_birth" json:"month_of_birth" validate:"required,lte=50"`
	YearOfBirth         int16     `db:"year_of_birth" json:"year_of_birth" validate:"required,lte=100"`
	AvatarUrl           string    `db:"avatar_url" json:"avatar_url" validate:"required,lte=100"`
	IsAutoRanking       bool      `db:"is_auto_ranking" json:"is_auto_ranking" validate:"required,lte=20"`
	IsAllowCustomerCare bool      `db:"is_allow_customer_care" json:"is_allow_customer_care" validate:"required,lte=50"`
	IsAllowMarketing    bool      `db:"is_allow_marketing" json:"is_allow_marketing" validate:"required,lte=50"`
	Note                string    `db:"note" json:"note" validate:"required,lte=100"`
	IsProspective       bool      `db:"is_prospective" json:"is_prospective" validate:"required,lte=100"`
}
