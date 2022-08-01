package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Requirement struct {
	ID                     uuid.UUID          `db:"id"`
	Title                  string             `db:"title"`
	Description            string             `db:"description"`
	CreatedBy              uuid.UUID          `db:"created_by"`
	CreatedAt              time.Time          `db:"created_at"`
	RequestingDepartmentID uuid.UUID          `db:"requesting_department_id"`
	ServiceDepartmentID    uuid.UUID          `db:"service_department_id"`
	RequirementTypeID      uuid.UUID          `db:"requirement_type_id"`
	RequirementSubTypeID   uuid.UUID          `db:"requirement_sub_type_id"`
	ModifiedBy             uuid.UUID          `db:"modified_by"`
	ModifiedAt             time.Time          `db:"modified_at"`
	ApprovedBy             uuid.UUID          `db:"approved_by"`
	ApprovedAt             time.Time          `db:"approved_at"`
	AcceptedBy             uuid.UUID          `db:"accepted_by"`
	AcceptedAt             time.Time          `db:"accepted_at"`
	DeclinedBy             uuid.UUID          `db:"declined_by"`
	DeclinedAt             time.Time          `db:"declined_at"`
	ProcessedBy            uuid.UUID          `db:"proccessed_by"`
	ProcessedAt            time.Time          `db:"proccessed_at"`
	AssignedBy             uuid.UUID          `db:"assigned_by"`
	AssignedTo             uuid.UUID          `db:"assigned_to"`
	AssignedAt             time.Time          `db:"assigned_at"`
	FinishedBy             uuid.UUID          `db:"finished_by"`
	FinishedAt             time.Time          `db:"finished_at"`
	Solved                 bool               `db:"solved"`
	Department             Department         `has_one:"department"`
	RequirementType        RequirementType    `has_one:"requirement_type"`
	RequirementSubType     RequirementSubType `has_one:"requirement_sub_type"`
	User                   User               `has_one:"user"`
}
