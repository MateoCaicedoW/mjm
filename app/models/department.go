package models

import (
	"time"

	"github.com/gofrs/uuid"
)

type Department struct {
	ID          uuid.UUID `form:"id" db:"id"`
	Name        string    `form:"name" db:"name"`
	Description string    `form:"description" db:"description"`
	CreatedAt   time.Time `form:"created_at" db:"created_at"`
	UpdatedAt   time.Time `form:"updated_at" db:"updated_at"`
}
