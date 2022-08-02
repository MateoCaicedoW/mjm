package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type RequirementSubType struct {
	ID                uuid.UUID        `json:"id" db:"id"`
	RequirementTypeID uuid.UUID        `json:"requirement_type_id" db:"requirement_type_id"`
<<<<<<< HEAD
	Name              string           `json:"name" db:"name"`
=======
	Name              string           `json:"name" db:"name" fako:"title"`
>>>>>>> requirements_crud
	CreatedAt         time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time        `json:"updated_at" db:"updated_at"`
	RequirementType   *RequirementType `belongs_to:"requirement_types"`
	Requirements      []Requirement    `has_many:"requirements"`
}
