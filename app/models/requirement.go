package models

import (
	"github.com/gofrs/uuid"
)

type Requirement struct {
	ID                     uuid.UUID      `db:"id"`
	Title                  string         `db:"title"`
	Description            string         `db:"description"`
	CreatedBy              uuid.UUID      `db:"created_by"`
	CreatedAt              uuid.Timestamp `db:"created_at"`
	RequestingDepartmentId uuid.UUID      `db:"requesting_department_id"`
	ServiceDepartmentId    uuid.UUID      `db:"service_department_id"`
	RequirementTypeId      uuid.UUID      `db:"requirement_type_id"`
	RequirementSubTypeId   uuid.UUID      `db:"requirement_sub_type_id"`
	ModifiedBy             uuid.UUID      `db:"modified_by"`
	ModifiedAt             uuid.Timestamp `db:"modified_at"`
	AprovedBy              uuid.UUID      `db:"aproved_by"`
	AprovedAt              uuid.Timestamp `db:"aproved_at"`
	DeclinedBy             uuid.UUID      `db:"declined_by"`
	DeclinedAt             uuid.Timestamp `db:"declined_at"`
	ProcessedBy            uuid.UUID      `db:"processed_by"`
	ProcessedAt            uuid.Timestamp `db:"processed_at"`
	AssignedBy             uuid.UUID      `db:"assigned_by"`
	AssignedTo             uuid.UUID      `db:"assigned_to"`
	AssignedAt             uuid.Timestamp `db:"assigned_at"`
	FinishedBy             uuid.UUID      `db:"finished_by"`
	FinishedAt             uuid.Timestamp `db:"finished_at"`
	Solved                 bool           `db:"solved"`
}
