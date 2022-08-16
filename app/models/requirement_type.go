package models

import (
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// RequirementType model struct.
type RequirementType struct {
	ID              uuid.UUID     `db:"id"`
	DepartmentID    uuid.UUID     `db:"department_id"`
	Name            string        `db:"name" fako:"brand"`
	Description     string        `db:"description"`
	CreatedByUserID uuid.UUID     `db:"user_id"`
	CreatedAt       time.Time     `db:"created_at"`
	UpdatedAt       time.Time     `db:"updated_at"`
	Department      *Department   `belongs_to:"departments"`
	CreatedByUser   *User         `belongs_to:"users"`
	Requirements    []Requirement `has_many:"requirements"`
}

type RequirementTypes []RequirementType

func (r *RequirementType) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{
			Field:   r.Name,
			Name:    "Name",
			Message: "Name is required.",
		},
		&validators.StringIsPresent{
			Field:   r.Description,
			Name:    "Description",
			Message: "Description is required.",
		},
	), nil

}
func (r RequirementTypes) Map() map[string]uuid.UUID {
	requirementTypesMap := map[string]uuid.UUID{}
	for _, v := range r {
		requirementTypesMap[v.Name] = v.ID
	}
	return requirementTypesMap
}
