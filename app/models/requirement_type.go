package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type RequirementType struct {
	ID           uuid.UUID `db:"id"`
	DepartmentId uuid.UUID `db:"department_id"`
	Name         string    `db:"name"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}
