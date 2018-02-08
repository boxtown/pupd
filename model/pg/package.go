package pg

import (
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	// Only import pq if using pg model package
	_ "github.com/lib/pq"
)

// IDGenerator is an interface for a universally unique
// ID generator
type IDGenerator interface {
	Generate() (string, error)
}

// UUIDV4Generator is an IDGenerator implementation
// that generates v4 UUIDs
type UUIDV4Generator struct{}

// Generate generates a v4 UUID string or returns an error
// if there was an issue during generation
func (gen UUIDV4Generator) Generate() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// DataSource represents a PostgreSQL data source
// to be used with `pg` stores
type DataSource struct {
	*sqlx.DB

	// Whether or not to print debug messages
	Debug bool
}

// NewDataSource creates a new `DataSource`
// connected to the PostgreSQL database indicated
// by the supplied connection string. An error is returned
// if there was an error connecting to the database
// Connection string examples can be found here:
// https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters
func NewDataSource(connStr string) (*DataSource, error) {
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &DataSource{DB: db}, nil
}

// Transaction executes the given handler within a transaction.
// If an error is returned by the handler or occurs during execution
// of the `handler` function, there will be an attempt to rollback the
// transaction.
func (source DataSource) Transaction(handler func(sqlx.Ext) error) error {
	tx, err := source.Beginx()
	if err != nil {
		return err
	}
	err = handler(tx)
	if err == nil {
		err = tx.Commit()
	}
	if err != nil {
		if err := tx.Rollback(); err != nil && source.Debug {
			log.Println(err)
		}
	}
	return err
}
