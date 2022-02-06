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

// Booking represents a customer's booking for a certain class
type Booking struct {
	ID       int
	Date     time.Time `json:"date" db:"date"`
	Customer string    `json:"customer" db:"customer"`
	Class    string    `json:"class" db:"class"`
}
