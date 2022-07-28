package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type RequirementSubTypes struct {
	ID                uuid.UUID `json:"id" db:"id"`
	RequirementTypeID uuid.UUID `json:"requirement_type_id" db:"requirement_type_id"`
	Name              string    `json:"name" db:"name"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}
