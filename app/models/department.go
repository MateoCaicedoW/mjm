package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Department struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name" fako:"first_name"`
	Description string    `db:"description" fako:"paragraph"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
