package storage

import (
	"flag"
	"os"
	"testing"

	"github.com/masci/go-rest-playground/models"
)

var storageType = flag.String("storage-type", "volatile", "type of storage to test: [volatile|sqlite]")

var getStorage func() Storage

func TestMain(m *testing.M) {
	if *storageType == "volatile" {
		getStorage = func() Storage {
			return NewVolatileStorage()
		}
	} else if *storageType == "sqlite" {
		getStorage = func() Storage {
			return NewSqliteStorage(":memory:")
		}
	} else {
		os.Exit(1)
	}

	os.Exit(m.Run())
}

func TestAddClass(t *testing.T) {
	s := getStorage()

	// count the fixtures
	list, _ := s.GetClasses()
	size := len(list)

	// add a class and check we have one item more
	s.AddClass(&models.Class{
		Name: "Foo",
	})
	if list, _ := s.GetClasses(); len(list) != size+1 {
		t.Errorf("got %d, want %d", len(list), size+1)
	}
}

func TestGetClass(t *testing.T) {
	s := getStorage()

	c, err := s.GetClass("PI0001")

	if err != nil {
		t.Errorf("got %s", err)
	}
	if c.Name != "Pilates" {
		t.Errorf("got %s, want %s", c.Name, "Pilates")
	}
}

func TestUpdateClass(t *testing.T) {
	s := getStorage()

	// wrong input
	c := &models.Class{}
	err := s.UpdateClass("wrong id!", c)
	if err == nil {
		t.Errorf("got nil, want error")
	}
	// input ok
	c, _ = s.GetClass("PI0001")
	newName := "Pilates Plus"
	c.Name = newName
	err = s.UpdateClass("PI0001", c)
	if err != nil {
		t.Errorf("got %s", err)
	}
	// get it again
	c, _ = s.GetClass("PI0001")
	if c.Name != newName {
		t.Errorf("got %s, want %s", c.Name, newName)
	}
}

func TestDeleteClass(t *testing.T) {
	s := getStorage()

	err := s.DeleteClass("PI0001")
	if err != nil {
		t.Errorf("got %s", err)
	}
	// ensure isn't there anymore
	_, err = s.GetClass("PI0001")
	if err == nil {
		t.Errorf("got nil, want error")
	}
}

func TestAddBooking(t *testing.T) {
	s := getStorage()

	// missing class id
	_, err := s.AddBooking(&models.Booking{})
	if err == nil {
		t.Errorf("got nil, want error")
	}
	// missing booking date
	_, err = s.AddBooking(&models.Booking{
		Class: "PI0001",
	})
	if err == nil {
		t.Errorf("got nil, want error")
	}
	// input ok
	id, err := s.AddBooking(&models.Booking{
		Class: "PI0001",
		Date:  createTime("2020-01-31"),
	})
	if err != nil {
		t.Errorf("got %s", err)
	}
	if id != 1 {
		t.Errorf("got %d, want %d", id, 1)
	}
}

func TestGetBooking(t *testing.T) {
	s := getStorage()

	// wrong id
	_, err := s.GetBooking(-1)
	if err == nil {
		t.Errorf("got nil, want error")
	}

	// add a valid booking
	s.AddBooking(&models.Booking{
		Class: "PI0001",
		Date:  createTime("2020-01-31"),
	})

	// input ok
	b, err := s.GetBooking(1)
	if err != nil {
		t.Errorf("got %s", err)
	}
	if b.Class != "PI0001" {
		t.Errorf("got %s, want %s", b.Class, "PI0001")
	}
}

func TestGetBookings(t *testing.T) {
	s := getStorage()

	// add valid bookings
	s.AddBooking(&models.Booking{
		Class: "PI0001",
		Date:  createTime("2020-01-31"),
	})
	s.AddBooking(&models.Booking{
		Class: "DA0001",
		Date:  createTime("2020-01-31"),
	})

	bookings, err := s.GetBookings()
	if err != nil {
		t.Errorf("got %s", err)
	}
	if len(bookings) != 2 {
		t.Errorf("got %d, want %d", len(bookings), 2)
	}
}

func TestUpdateBooking(t *testing.T) {
	s := getStorage()

	// test invalid input
	err := s.UpdateBooking(-1, &models.Booking{})
	if err == nil {
		t.Errorf("got nil, want error")
	}

	// add valid bookings
	b := &models.Booking{
		Customer: "Foo",
		Class:    "PI0001",
		Date:     createTime("2020-01-31"),
	}
	id, _ := s.AddBooking(b)

	// update the Customer field
	b.Customer = "Bar"
	err = s.UpdateBooking(id, b)
	if err != nil {
		t.Errorf("got %s", err)
	}

	// reload to assert record was updated
	newb, _ := s.GetBooking(1)
	if newb.Customer != b.Customer {
		t.Errorf("got %s, want %s", newb.Customer, b.Customer)
	}
}

func TestDeleteBooking(t *testing.T) {
	s := getStorage()

	// add valid bookings
	b := &models.Booking{
		Customer: "Foo",
		Class:    "PI0001",
		Date:     createTime("2020-01-31"),
	}
	id, _ := s.AddBooking(b)

	err := s.DeleteBooking(id)
	if err != nil {
		t.Errorf("got %s", err)
	}
}

func TestClose(t *testing.T) {
	s := getStorage()

	err := s.Close()
	if err != nil {
		t.Errorf("got %s", err)
	}
}
