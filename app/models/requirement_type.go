package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type RequirementType struct {
<<<<<<< HEAD
	ID           uuid.UUID   `db:"id"`
	DepartmentID uuid.UUID   `db:"department_id"`
	Name         string      `db:"name"`
	CreatedAt    time.Time   `db:"created_at"`
	UpdatedAt    time.Time   `db:"updated_at"`
	Department   *Department `belongs_to:"departments"`
	Requirements []Requirement
=======
	ID           uuid.UUID     `db:"id"`
	DepartmentID uuid.UUID     `db:"department_id"`
	Name         string        `db:"name" fako:"title"`
	CreatedAt    time.Time     `db:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at"`
	Department   *Department   `belongs_to:"departments"`
	Requirements []Requirement `has_many:"requirements"`
>>>>>>> requirements_crud
}
