package storage

import (
	"fmt"

	"github.com/go-rest-playground/models"
)

// VolatileStorage implements a trivial in-memory storage for the
// Storage interface using maps.
type VolatileStorage struct {
	classes map[string]*models.Class
}

// NewVolatileStorage creates the data in memory and loads the initial fixtures.
func NewVolatileStorage() Storage {
	c := map[string]*models.Class{}
	for _, item := range classes {
		c[item.ID] = item
	}

	return &VolatileStorage{classes: c}
}

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
	for _, c := range s.classes {
		if c.ID == ID {
			return c, nil
		}
	}

	return nil, fmt.Errorf("no Class found with id '%s'", ID)
}

func (s *VolatileStorage) UpdateClass(ID string, c *models.Class) error {
	for i, c := range s.classes {
		if c.ID == ID {
			s.classes[i] = c
		}
	}

	return nil
}

func (s *VolatileStorage) DeleteClass(ID string) error {
	delete(s.classes, ID)
	return nil
}

func (s *VolatileStorage) Close() error {
	/* noop */
	return nil
}
