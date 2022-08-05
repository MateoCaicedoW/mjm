package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type RequirementSubType struct {
	ID                uuid.UUID        `json:"id" db:"id"`
	RequirementTypeID uuid.UUID        `json:"requirement_type_id" db:"requirement_type_id"`
	Name              string           `json:"name" db:"name"`
	CreatedAt         time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at" db:"updated_at"`
	RequirementType   *RequirementType `belongs_to:"requirement_types"`
	Requirements      []Requirement    `has_many:"requirements"`
}

type RequirementSubTypes []RequirementSubType

func (rs RequirementSubTypes) Map() map[string]uuid.UUID {
	subtypesOptions := map[string]uuid.UUID{}
	subtypesOptions["Select  a Subtype"] = uuid.Nil

	for _, e := range rs {
		subtypesOptions[e.Name] = e.ID
	}

	return subtypesOptions
}
