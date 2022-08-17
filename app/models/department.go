package models

import (
	"regexp"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

type Department struct {
	ID                uuid.UUID        `db:"id"`
	Name              string           `db:"name" fako:"first_name"`
	Description       string           `db:"description" fako:"paragraph"`
	CreatedAt         time.Time        `db:"created_at"`
	UpdatedAt         time.Time        `db:"updated_at"`
	Users             []User           `has_many:"users"`
	RequirementsTypes RequirementTypes `many_to_many:"area_requirement_types"`
	RequirementsType  map[string]bool  `db:"-" form:"requirements_type"`
}

type Departments []Department

func (d Departments) Map() map[string]uuid.UUID {
	departmentsMap := map[string]uuid.UUID{}

	for _, v := range d {
		departmentsMap[v.Name] = v.ID
	}
	return departmentsMap
}

func (d *Department) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{
			Field:   d.Name,
			Name:    "Name",
			Message: "Name is required.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if d.Name != "" && !regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString(d.Name) {
					return false
				}
				return true
			},
			Name:    "Name",
			Message: "%s Name must be letters only.",
		},
		&validators.StringIsPresent{
			Field:   d.Description,
			Name:    "Description",
			Message: "Description is required.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if d.Name != "" && len(d.Name) > 255 {
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
				if len(d.RequirementsType) == 0 {
					return len(d.RequirementsType) != 0
				}
				return true
			},
			Field:   "",
			Name:    "RequirementsType",
			Message: "%s Some types must be selected.",
		},
	), nil

}
