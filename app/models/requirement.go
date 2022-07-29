package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Requirement struct {
	ID                     uuid.UUID `db:"id"`
	Title                  string    `db:"title"`
	Description            string    `db:"description"`
	CreatedBy              uuid.UUID `db:"created_by"`
	CreatedAt              time.Time `db:"created_at"`
	RequestingDepartmentId uuid.UUID `db:"requesting_department_id"`
	ServiceDepartmentId    uuid.UUID `db:"service_department_id"`
	RequirementTypeId      uuid.UUID `db:"requirement_type_id"`
	RequirementSubTypeId   uuid.UUID `db:"requirement_sub_type_id"`
	ModifiedBy             uuid.UUID `db:"modified_by"`
	ModifiedAt             time.Time `db:"modified_at"`
	AprovedBy              uuid.UUID `db:"aproved_by"`
	AprovedAt              time.Time `db:"aproved_at"`
	AcceptedBy             uuid.UUID `db:"accepted_by"`
	AcceptedAt             time.Time `db:"accepted_at"`
	DeclinedBy             uuid.UUID `db:"declined_by"`
	DeclinedAt             time.Time `db:"declined_at"`
	ProcessedBy            uuid.UUID `db:"proccessed_by"`
	ProcessedAt            time.Time `db:"proccessed_at"`
	AssignedBy             uuid.UUID `db:"assigned_by"`
	AssignedTo             uuid.UUID `db:"assigned_to"`
	AssignedAt             time.Time `db:"assigned_at"`
	FinishedBy             uuid.UUID `db:"finished_by"`
	FinishedAt             time.Time `db:"finished_at"`
	Solved                 bool      `db:"solved"`
}
