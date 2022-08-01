package models

import (
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gofrs/uuid"
)

type Requirement struct {
	ID                     uuid.UUID           `db:"id"`
	Title                  string              `db:"title" fako:"job_title"`
	Description            string              `db:"description" fako:"sentence"`
	CreatedByUserID        uuid.UUID           `db:"created_by"`
	CreatedAt              time.Time           `db:"created_at"`
	RequestingDepartmentID uuid.UUID           `db:"requesting_department_id"`
	ServiceDepartmentID    uuid.UUID           `db:"service_department_id"`
	RequirementSubTypeID   uuid.UUID           `db:"requirement_sub_type_id"`
	RequirementTypeID      uuid.UUID           `db:"requirement_type_id"`
	ModifiedByUserID       nulls.UUID          `db:"modified_by"`
	ModifiedAt             time.Time           `db:"modified_at"`
	ApprovedByUserID       nulls.UUID          `db:"approved_by"`
	ApprovedAt             time.Time           `db:"approved_at"`
	AcceptedByUserID       nulls.UUID          `db:"accepted_by"`
	AcceptedAt             time.Time           `db:"accepted_at"`
	DeclinedByUserID       nulls.UUID          `db:"declined_by"`
	DeclinedAt             time.Time           `db:"declined_at"`
	ProcessedByUserID      nulls.UUID          `db:"processed_by"`
	ProcessedAt            time.Time           `db:"processed_at"`
	AssignedByUserID       nulls.UUID          `db:"assigned_by"`
	AssignedToUserID       nulls.UUID          `db:"assigned_to"`
	AssignedAt             time.Time           `db:"assigned_at"`
	FinishedByUserID       nulls.UUID          `db:"finished_by"`
	FinishedAt             time.Time           `db:"finished_at"`
	UpdatedAt              time.Time           `db:"updated_at"`
	Solved                 bool                `db:"solved"`
	CreatedByUser          *User               `belongs_to:"users"`
	RequestingDepartment   *Department         `belongs_to:"departments"`
	ServiceDepartment      *Department         `belongs_to:"departments"`
	RequirementSubType     *RequirementSubType `belongs_to:"requirement_sub_types"`
	RequirementType        *RequirementType    `belongs_to:"requirement_types"`
	ModifiedByUser         *User               `belongs_to:"users"`
	ApprovedByUser         *User               `belongs_to:"users"`
	AcceptedByUser         *User               `belongs_to:"users"`
	DeclinedByUser         *User               `belongs_to:"users"`
	ProcessedByUser        *User               `belongs_to:"users"`
	AssignedToUser         *User               `belongs_to:"users"`
	AssignedByUser         *User               `belongs_to:"users"`
	FinishedByUser         *User               `belongs_to:"users"`
}

var Requirements []Requirement
