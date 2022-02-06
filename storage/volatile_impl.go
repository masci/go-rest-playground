package storage

import (
	"fmt"

	"github.com/masci/go-rest-playground/models"
)

// VolatileStorage implements a trivial in-memory storage for the
// Storage interface using maps.
type VolatileStorage struct {
	classes         map[string]*models.Class
	bookings        map[int]*models.Booking
	last_booking_id int
}

// NewVolatileStorage creates the data in memory and loads the initial fixtures.
func NewVolatileStorage() Storage {
	c := map[string]*models.Class{}
	for _, item := range classes {
		c[item.ID] = item
	}

	return &VolatileStorage{
		classes:  c,
		bookings: map[int]*models.Booking{},
	}
}

/*
	Class management functions
*/

func (s *VolatileStorage) AddClass(c *models.Class) (string, error) {
	c.ID = makeID(c.Name)
	s.classes[c.ID] = c
	return c.ID, nil
}

func (s *VolatileStorage) GetClasses() ([]*models.Class, error) {
	retVal := []*models.Class{}
	for _, val := range s.classes {
		retVal = append(retVal, val)
	}

	return retVal, nil
}

func (s *VolatileStorage) GetClass(ID string) (*models.Class, error) {
	val, ok := s.classes[ID]
	if ok {
		return val, nil
	}

	return nil, fmt.Errorf("no Class found with id '%s'", ID)
}

func (s *VolatileStorage) UpdateClass(ID string, c *models.Class) error {
	_, ok := s.classes[ID]
	if ok {
		s.classes[ID] = c
		return nil
	}

	return fmt.Errorf("no Class found with id '%s'", ID)
}

func (s *VolatileStorage) DeleteClass(ID string) error {
	delete(s.classes, ID)
	return nil
}

/*
	Booking management functions
*/

func (s *VolatileStorage) AddBooking(b *models.Booking) (int, error) {
	// check wether the booking is valid
	class, err := s.GetClass(b.Class)
	if err != nil {
		return -1, err
	}
	if !canBook(b, class) {
		return -1, fmt.Errorf("class %s is not available at %s", class.Name, b.Date)
	}

	// proceed with booking creation
	s.last_booking_id++
	b.ID = s.last_booking_id
	s.bookings[b.ID] = b
	return b.ID, nil
}

func (s *VolatileStorage) GetBookings() ([]*models.Booking, error) {
	retVal := []*models.Booking{}
	for _, val := range s.bookings {
		retVal = append(retVal, val)
	}

	return retVal, nil
}

func (s *VolatileStorage) GetBooking(ID int) (*models.Booking, error) {
	val, ok := s.bookings[ID]
	if ok {
		return val, nil
	}

	return nil, fmt.Errorf("no Booking found with id '%d'", ID)

}

func (s *VolatileStorage) UpdateBooking(ID int, booking *models.Booking) error {
	_, ok := s.bookings[ID]
	if ok {
		s.bookings[ID] = booking
		return nil
	}

	return fmt.Errorf("no Booking found with id '%d'", ID)
}

func (s *VolatileStorage) DeleteBooking(ID int) error {
	delete(s.bookings, ID)
	return nil
}

/*
	Others
*/

func (s *VolatileStorage) Close() error {
	/* noop */
	return nil
}
