// The models package implements the types and logic needed to represent
// our data model.
package models

import (
	"time"
)

// Class represents a class in a gym or studio
type Class struct {
	ID        string
	Name      string    `json:"name" db:"name"`
	StartDate time.Time `json:"start_date" db:"start_date"`
	EndDate   time.Time `json:"end_date" db:"end_date"`
	Capacity  int       `json:"capacity" db:"capacity"`
}
