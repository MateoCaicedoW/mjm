package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// RequirementType model struct.
type RequirementType struct {
	ID                  uuid.UUID             `db:"id"`
	Name                string                `db:"name" fako:"brand"`
	CreatedAt           time.Time             `db:"created_at"`
	UpdatedAt           time.Time             `db:"updated_at"`
	Requirements        []Requirement         `has_many:"requirements"`
	AreaRequirementType []AreaRequirementType `has_many:"areas_requirements_types"`
}

type RequirementTypes []RequirementType

func (rt RequirementTypes) Map() map[uuid.UUID]string {

	requirementTypeMap := map[uuid.UUID]string{}

	for _, e := range rt {
		requirementTypeMap[e.ID] = e.Name
	}

	return requirementTypeMap
}
