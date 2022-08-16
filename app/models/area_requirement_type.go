package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type AreaRequirementType struct {
	ID                uuid.UUID        `db:"id"`
	DepartmentID      uuid.UUID        `db:"department_id"`
	RequirementTypeID uuid.UUID        `db:"requirement_type_id"`
	Department        *Department      `belongs_to:"departments"`
	RequirementType   *RequirementType `belongs_to:"requirement_types"`
	CreatedAt         time.Time        `db:"created_at"`
	UpdatedAt         time.Time        `db:"updated_at"`
}
