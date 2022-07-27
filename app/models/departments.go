package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Departments struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	RequirementType       *Requirement_type     `belongs_to:"requirement_types"`
}
