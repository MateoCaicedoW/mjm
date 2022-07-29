package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Department struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type Departments []Department

func (d Departments) Map() map[string]uuid.UUID {
	departmentsMap := map[string]uuid.UUID{}
	for _, v := range d {
		departmentsMap[v.Name] = v.ID
	}
	return departmentsMap
}
