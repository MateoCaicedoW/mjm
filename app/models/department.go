package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Department struct {
	ID           uuid.UUID     `db:"id"`
	Name         string        `db:"name" fako:"first_name"`
	Description  string        `db:"description" fako:"paragraph"`
	CreatedAt    time.Time     `db:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at"`
	Users        []User        `has_many:"users"`
	Requirements []Requirement `has_many:"requirements"`
}

type Departments []Department

func (d Departments) Map() map[string]uuid.UUID {
	departmentsMap := map[string]uuid.UUID{}
	
	for _, v := range d {
		departmentsMap[v.Name] = v.ID
	}
	return departmentsMap
}

