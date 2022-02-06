package storage

import (
	"github.com/go-rest-playground/models"
	_ "github.com/mattn/go-sqlite3"
)

// Storage is the public API of our storage system. In this example
// we provide two concrete implementations of this interface:
// VolatileStorage and SqliteStorage
type Storage interface {
	// Class
	AddClass(*models.Class) (string, error)
	GetClasses() ([]*models.Class, error)
	GetClass(ID string) (*models.Class, error)
	UpdateClass(ID string, class *models.Class) error
	DeleteClass(ID string) error

	// Booking
	AddBooking(*models.Booking) (int, error)
	GetBookings() ([]*models.Booking, error)
	GetBooking(ID int) (*models.Booking, error)
	UpdateBooking(ID int, booking *models.Booking) error
	DeleteBooking(ID int) error

	// Others
	Close() error
}
