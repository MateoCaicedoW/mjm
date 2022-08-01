package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Requirement struct {
	ID                     uuid.UUID           `db:"id"`
	Title                  string              `db:"title"`
	Description            string              `db:"description"`
	CreatedBy              uuid.UUID           `db:"created_by"`
	CreatedAt              time.Time           `db:"created_at"`
	RequestingDepartmentId uuid.UUID           `db:"requesting_department_id"`
	ServiceDepartmentId    uuid.UUID           `db:"service_department_id"`
	RequirementTypeId      uuid.UUID           `db:"requirement_type_id"`
	RequirementSubTypeId   uuid.UUID           `db:"requirement_sub_type_id"`
	ModifiedBy             uuid.UUID           `db:"modified_by"`
	ModifiedAt             time.Time           `db:"modified_at"`
	ApprovedBy             uuid.UUID           `db:"approved_by"`
	ApprovedAt             time.Time           `db:"approved_at"`
	AcceptedBy             uuid.UUID           `db:"accepted_by"`
	AcceptedAt             time.Time           `db:"accepted_at"`
	DeclinedBy             uuid.UUID           `db:"declined_by"`
	DeclinedAt             time.Time           `db:"declined_at"`
	ProcessedBy            uuid.UUID           `db:"proccessed_by"`
	ProcessedAt            time.Time           `db:"proccessed_at"`
	AssignedBy             uuid.UUID           `db:"assigned_by"`
	AssignedTo             uuid.UUID           `db:"assigned_to"`
	AssignedAt             time.Time           `db:"assigned_at"`
	FinishedBy             uuid.UUID           `db:"finished_by"`
	FinishedAt             time.Time           `db:"finished_at"`
	UpdatedAt              time.Time           `db:"updated_at"`
	Solved                 bool                `db:"solved"`
	CreatedByUser          *User               `has_one:"user"`
	Department             *Department         `has_one:"department"`
	ServiceDepartment      *Department         `has_one:"department"`
	RequirementType        *RequirementType    `has_one:"requirement_type"`
	RequirementSubType     *RequirementSubType `has_one:"requirement_sub_type"`
	ModifiedByUser         *User               `has_one:"user"`
	ApprovedByUser         *User               `has_one:"user"`
	AcceptedByUser         *User               `has_one:"user"`
	DeclinedByUser         *User               `has_one:"user"`
	ProcessedByUser        *User               `has_one:"user"`
	AssignedByUser         *User               `has_one:"user"`
	FinishedByUser         *User               `has_one:"user"`
}

var Requirements []Requirement
