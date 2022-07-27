package models

import "github.com/gofrs/uuid"

type RequerimentType struct {
	ID           uuid.UUID `db:"id"`
	DepartmentId uuid.UUID `db:"department_id"`
	Name         string    `db:"name"`
}