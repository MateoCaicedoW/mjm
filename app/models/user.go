package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID           uuid.UUID     `db:"id" `
	FirstName    string        `db:"first_name" fako:"first_name" `
	LastName     string        `db:"last_name" fako:"last_name"`
	DNI          string        `db:"dni" fako:"phone"`
	EmailAddress string        `db:"email_address"  fako:"email_address"`
	PhoneNumber  string        `db:"phone_number" fako:"phone" `
	CreatedAt    time.Time     `db:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at"`
	DepartmentID uuid.UUID     `db:"department_id"`
	Department   *Department   `belongs_to:"departments"`
	Requirements []Requirement `has_many:"requirements"`
}
