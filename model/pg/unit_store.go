package pg

import (
	"github.com/boxtown/pupd/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// UnitStore implements model.UnitStore
// using PostgreSQL
type UnitStore struct {
	source sqlx.Ext
}

// NewUnitStore returns a PostgreSQL-backed implementation
// of model.UnitStore
func NewUnitStore(source sqlx.Ext) model.UnitStore {
	return &UnitStore{source: source}
}

// Create attempts to create a record for the given Unit.
// A v4 UUID will be assigned to the Unit and returned by
// this method
func (store UnitStore) Create(unit *model.Unit) (string, error) {
	raw, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	id := raw.String()

	if _, err = store.source.Exec(
		"INSERT INTO core.units (unit_id, name) VALUES ($1, $2)",
		id,
		unit.Name,
	); err != nil {
		return "", err
	}
	return id, nil
}

// Get attempts to retrieve a Unit from storage
// by its ID
func (store UnitStore) Get(id string) (*model.Unit, error) {
	row := store.source.QueryRowx(
		"SELECT unit_id, name FROM core.units WHERE unit_id=$1",
		id,
	)
	unit := model.Unit{}
	if err := row.Scan(&unit.ID, &unit.Name); err != nil {
		return nil, err
	}
	return &unit, nil
}

// List lists all Units from storage
func (store UnitStore) List() ([]model.Unit, error) {
	rows, err := store.source.Query("SELECT unit_id, name FROM core.units")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	units := []model.Unit{}
	for rows.Next() {
		unit := model.Unit{}
		if err := rows.Scan(&unit.ID, &unit.Name); err != nil {
			return nil, err
		}
		units = append(units, unit)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return units, nil
}

// Update attempts to update a Unit's record in storage
func (store UnitStore) Update(unit *model.Unit) error {
	_, err := store.source.Exec(
		"UPDATE core.units SET name=$1 WHERE unit_id=$2",
		unit.Name,
		unit.ID,
	)
	return err
}

// Delete attempts to delete a Unit's record in storage
// by its ID
func (store UnitStore) Delete(id string) error {
	_, err := store.source.Exec(
		"DELETE FROM core.units WHERE unit_id=$1",
		id,
	)
	return err
}
