package model

import "github.com/jmoiron/sqlx"

// DataSource represents a transactional data source
// interface
type DataSource interface {
	sqlx.Ext
	Transaction(handler func(sqlx.Ext) error) error
}
