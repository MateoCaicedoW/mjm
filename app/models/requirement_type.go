package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// RequirementType model struct.
type RequirementType struct {
	ID           uuid.UUID     `db:"id"`
	DepartmentID uuid.UUID     `db:"department_id"`
	Name         string        `db:"name" fako:"brand"`
	CreatedAt    time.Time     `db:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at"`
	Department   *Department   `belongs_to:"departments"`
	Requirements []Requirement `has_many:"requirements"`
}

type RequirementTypes []RequirementType
