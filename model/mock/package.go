// Code generated. This is hack to turn off linter for mocks. DO NOT EDIT.
// This will hopefully actually be code generated once github.com/golang/mock
// gets their shit together and merges in #28

package mock

import "github.com/jmoiron/sqlx"

type MockDataSource struct {
	sqlx.Ext
	TransactionFn func(handler func(sqlx.Ext) error) error
}

func (source MockDataSource) Transaction(handler func(sqlx.Ext) error) error {
	return source.TransactionFn(handler)
}
