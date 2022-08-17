package models

import (
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

type Users []User
type User struct {
	ID           uuid.UUID     `db:"id" `
	FirstName    string        `db:"first_name" `
	LastName     string        `db:"last_name" `
	DNI          string        `db:"dni" `
	EmailAddress string        `db:"email_address" `
	PhoneNumber  string        `db:"phone_number" `
	CreatedAt    time.Time     `db:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at"`
	DepartmentID uuid.UUID     `db:"department_id"`
	Department   *Department   `belongs_to:"departments"`
	Requirements []Requirement `has_many:"requirements"`
}

func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{
			Field:   u.FirstName,
			Name:    "FirstName",
			Message: "First name is required.",
		},
		&validators.StringIsPresent{
			Field:   u.LastName,
			Name:    "LastName",
			Message: "Last name is required.",
		},&validators.StringIsPresent{
			Field:   u.DNI,
			Name:    "DNI",
			Message: "DNI is required.",
		},&validators.StringIsPresent{
			Field:   u.EmailAddress,
			Name:    "EmailAddress",
			Message: "Email address is required.",
		},
		&validators.StringIsPresent{
			Field:   u.PhoneNumber,
			Name:    "PhoneNumber",
			Message: "Phone number is required.",
		},
	), nil

}

func (u Users) Map() map[string]uuid.UUID {
	departmentsMap := map[string]uuid.UUID{}
	for _, v := range u {
		departmentsMap[v.FirstName] = v.ID
	}
	return departmentsMap
}
