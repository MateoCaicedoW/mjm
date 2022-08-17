package models

import (
	"regexp"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// RequirementType model struct.
type RequirementType struct {
	ID              uuid.UUID     `db:"id"`
	Name            string        `db:"name" fako:"brand"`
	Description     string        `db:"description"`
	CreatedByUserID uuid.UUID     `db:"user_id"`
	CreatedAt       time.Time     `db:"created_at"`
	UpdatedAt       time.Time     `db:"updated_at"`
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
		&validators.FuncValidator{
			Fn: func() bool {
				if r.Name != "" && len(r.Name) > 255 {
					return false
				}
				return true
			},
			Field:   "",
			Name:    "Name",
			Message: "%s Name must be less than 255 characters.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if r.Description != "" && len(r.Description) > 255 {
					return false
				}
				return true
			},
			Field:   "",
			Name:    "Description",
			Message: "%s Description must be less than 255 characters.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if r.Name != "" && !regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString(r.Name) {
					return false
				}
				return true
			},
			Name:    "Name",
			Message: "%s Name must be letters only.",
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
