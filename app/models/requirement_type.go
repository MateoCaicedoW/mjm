package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// RequirementType model struct.
type RequirementType struct {
	ID           uuid.UUID     `db:"id"`
	Name         string        `db:"name" fako:"brand"`
	CreatedAt    time.Time     `db:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at"`
	Requirements []Requirement `has_many:"requirements"`
}

type RequirementTypes []RequirementType

func (rt RequirementTypes) Map() map[string]uuid.UUID {

	serviceArea := map[string]uuid.UUID{}

	for _, e := range rt {
		serviceArea[e.Name] = e.ID
	}

	return serviceArea
}
