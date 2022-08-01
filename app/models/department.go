package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Department struct {
	ID           uuid.UUID     `json:"id" db:"id"`
	Name         string        `json:"name" db:"name" fako:"product_name"`
	Description  string        `json:"description" db:"description" fako:"description"`
	CreatedAt    time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at" db:"updated_at"`
	Users        []User        `has_many:"users"`
	Requirements []Requirement `has_many:"requirements"`
}
