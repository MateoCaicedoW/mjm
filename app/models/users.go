package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID           uuid.UUID     `db:"id" `
	Email        string        `db:"email" `
	FirstName    string        `db:"first_name" `
	LastName     string        `db:"last_name" `
	DNI          string        `db:"dni" `
	EmailAddress string        `db:"email_address" `
	PhoneNumber  string        `db:"phone_number" `
	CreatedAt    time.Time     `db:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at"`
	Departments  []Departments `has_many:"deparments"`
}
