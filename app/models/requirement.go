package models

import (
	"regexp"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
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
	ModifiedAt             nulls.Time          `db:"modified_at"`
	ApprovedByUserID       nulls.UUID          `db:"approved_by"`
	ApprovedAt             nulls.Time          `db:"approved_at"`
	AcceptedByUserID       nulls.UUID          `db:"accepted_by"`
	AcceptedAt             nulls.Time          `db:"accepted_at"`
	DeclinedByUserID       nulls.UUID          `db:"declined_by"`
	DeclinedAt             nulls.Time          `db:"declined_at"`
	ProcessedByUserID      nulls.UUID          `db:"processed_by"`
	ProcessedAt            nulls.Time          `db:"processed_at"`
	AssignedByUserID       nulls.UUID          `db:"assigned_by"`
	AssignedToUserID       nulls.UUID          `db:"assigned_to"`
	AssignedAt             nulls.Time          `db:"assigned_at"`
	FinishedByUserID       nulls.UUID          `db:"finished_by"`
	FinishedAt             nulls.Time          `db:"finished_at"`
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

func (r *Requirement) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{
			Field:   r.Title,
			Name:    "Title",
			Message: "Title is required.",
		},
		&validators.StringIsPresent{
			Field:   r.Description,
			Name:    "Description",
			Message: "Description is required.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if r.Title != "" && len(r.Title) > 255 {
					return false
				}
				return true
			},
			Field:   "",
			Name:    "Title",
			Message: "%s Title must be less than 255 characters.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if r.Description != "" && len(r.Description) > 255 {
					return false
				}
				return true
			},
			Field:   "",
			Name:    "Description",
			Message: "%s Description must be less than 255 characters.",
		},
		&validators.FuncValidator{
			Fn: func() bool {
				if r.Title != "" && !regexp.MustCompile(`^[a-zA-Z ]+$`).MatchString(r.Title) {
					return false
				}
				return true
			},
			Name:    "Title",
			Message: "%s Title must be letters only.",
		},
		&validators.UUIDIsPresent{
			Field:   r.CreatedByUserID,
			Name:    "CreatedByUserID",
			Message: "User is required.",
		},
		&validators.UUIDIsPresent{
			Field:   r.RequestingDepartmentID,
			Name:    "RequestingDepartmentID",
			Message: "Area is required.",
		},
		&validators.UUIDIsPresent{
			Field:   r.ServiceDepartmentID,
			Name:    "ServiceDepartmentID",
			Message: "Service Area is required.",
		},
		&validators.UUIDIsPresent{
			Field:   r.RequirementTypeID,
			Name:    "RequirementTypeID",
			Message: "Type is required.",
		},
	), nil

}
