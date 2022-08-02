package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// RequirementType model struct.
type RequirementType struct {
	ID           uuid.UUID `form:"id" db:"id"`
	DepartmentID uuid.UUID `form:"department_id" db:"department_id"`
	Name         string    `form:"name" db:"name" fako:"domain_name"`
	CreatedAt    time.Time `form:"created_at" db:"created_at"`
	UpdatedAt    time.Time `form:"updated_at" db:"updated_at"`
}

type RequirementTypes []RequirementType
