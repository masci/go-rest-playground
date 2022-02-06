package storage

import (
	"github.com/go-rest-playground/models"
	"github.com/jmoiron/sqlx"
)

var schema = `
CREATE TABLE class (
	id TEXT PRIMARY KEY,
    name TEXT,
	start_date DATETIME,
	end_date DATETIME,
    capacity INTEGER
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

func (s *SqliteStorage) Close() error {
	return s.db.Close()
}

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
