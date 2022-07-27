package models

import "github.com/gofrs/uuid"

type Requirement_type struct {
	ID           uuid.UUID     `db:"id"`
	DepartmentId uuid.UUID     `db:"department_id"`
	Departments  []Departments `has_many:"deparments"`
}
