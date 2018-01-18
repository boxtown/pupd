package pg

import (
	// Only import pq if using pg model package
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

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
func (source DataSource) Transaction(handler func(*sqlx.Tx) error) error {
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
