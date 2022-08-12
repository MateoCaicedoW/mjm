package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type AreaRequirementType struct {
	ID                uuid.UUID        `db:"id"`
	DepartmentID      uuid.UUID        `db:"area_id"`
	RequirementTypeID uuid.UUID        `db:"requirement_type_id"`
	Department        *Department      `belongs_to:"department"`
	RequirementType   *RequirementType `belongs_to:"requirement_types"`
	CreatedAt         time.Time        `db:"created_at"`
	UpdatedAt         time.Time        `db:"updated_at"`
}
