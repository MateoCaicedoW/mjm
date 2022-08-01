package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Department struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name" fako:"job_name"`
	Description string    `db:"description" fako:"sentence"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
