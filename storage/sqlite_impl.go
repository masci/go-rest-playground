package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/masci/go-rest-playground/models"
)

var schema = `
CREATE TABLE class (
	id TEXT PRIMARY KEY,
    name TEXT,
	start_date DATETIME,
	end_date DATETIME,
    capacity INTEGER
);

CREATE TABLE booking (
	id INTEGER PRIMARY KEY,
	date DATETIME,
	customer TEXT,
	class TEXT
);
`

// SqliteStorage implements the Storage interface saving data
// in a SQLite database on disk.
type SqliteStorage struct {
	db *sqlx.DB
}

// NewSqliteStorage creates the database on file and loads the initial fixtures.
// The path to the database file is passed by the caller with the `path` parameter.
func NewSqliteStorage(path string) Storage {
	d := &SqliteStorage{
		db: sqlx.MustConnect("sqlite3", path),
	}
	d.db.Exec(schema)

	tx := d.db.MustBegin()
	for _, item := range classes {
		tx.NamedExec(
			"INSERT INTO class(id, name, start_date, end_date, capacity) VALUES (:id, :name, :start_date, :end_date, :capacity)",
			item,
		)
	}
	tx.Commit()

	return d
}

/*
	Class management functions
*/

func (s *SqliteStorage) AddClass(c *models.Class) (string, error) {
	c.ID = makeID(c.Name)

	_, err := s.db.NamedExec(
		"INSERT INTO class(id, name, start_date, end_date, capacity) VALUES (:id, :name, :start_date, :end_date, :capacity)",
		c,
	)

	return c.ID, err
}

func (s *SqliteStorage) GetClasses() ([]*models.Class, error) {
	classes := []*models.Class{}

	err := s.db.Select(&classes, `SELECT * FROM class`)

	return classes, err
}

func (s *SqliteStorage) GetClass(ID string) (*models.Class, error) {
	c := models.Class{}
	err := s.db.Get(&c, "SELECT * FROM class WHERE id=$1", ID)

	return &c, err
}

func (s *SqliteStorage) UpdateClass(ID string, c *models.Class) error {

	_, err := s.db.NamedExec(
		"Update class SET name=:name, start_date=:start_date, end_date=:end_date, capacity=:capacity WHERE id=:id",
		c,
	)

	return err
}

func (s *SqliteStorage) DeleteClass(ID string) error {
	_, err := s.db.Exec("DELETE from class WHERE id=$1", ID)
	return err
}

/*
	Class management functions
*/

func (s *SqliteStorage) AddBooking(b *models.Booking) (int, error) {
	err := s.db.Get(&b.ID, "SELECT IFNULL( MAX(id), 0 ) from booking;") // this strategy won't reuse deleted ids
	if err != nil {
		return 0, err
	}
	b.ID++

	_, err = s.db.NamedExec(
		"INSERT INTO booking(id, date, customer, class) VALUES (:id, :date, :customer, :class)",
		b,
	)

	return b.ID, err
}

func (s *SqliteStorage) GetBookings() ([]*models.Booking, error) {
	bookings := []*models.Booking{}

	err := s.db.Select(&classes, `SELECT * FROM bookings`)

	return bookings, err
}

func (s *SqliteStorage) GetBooking(ID int) (*models.Booking, error) {
	b := models.Booking{}
	err := s.db.Get(&b, "SELECT * FROM booking WHERE id=$1", ID)

	return &b, err
}

func (s *SqliteStorage) UpdateBooking(ID int, c *models.Booking) error {

	_, err := s.db.NamedExec(
		"Update booking SET date=:date, customer=:customer, class=:class WHERE id=:id",
		c,
	)

	return err
}

func (s *SqliteStorage) DeleteBooking(ID int) error {
	_, err := s.db.Exec("DELETE from booking WHERE id=$1", ID)
	return err
}

/*
	Others
*/

func (s *SqliteStorage) Close() error {
	return s.db.Close()
}
