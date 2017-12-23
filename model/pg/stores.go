package pg

import (
	"github.com/boxtown/pupd/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"

	// Only import pq if using pg model package
	_ "github.com/lib/pq"
)

type MovementStore struct {
	source sqlx.Ext
}

func NewMovementStore(source sqlx.Ext) model.MovementStore {
	return &MovementStore{source: source}
}

func (store MovementStore) Create(movement *model.Movement) (string, error) {
	if len(movement.ID) == 0 {
		id, err := uuid.NewRandom()
		if err != nil {
			return "", err
		}
		movement.ID = id.String()
	}
	_, err := store.source.Exec(
		"INSERT INTO core.movements (movement_id, name) VALUES ($1, $2)",
		movement.ID,
		movement.Name,
	)
	if err != nil {
		return "", err
	}
	return movement.ID, nil
}

func (store MovementStore) Get(id string) (*model.Movement, error) {
	row := store.source.QueryRowx(
		"SELECT movement_id, name FROM core.movements WHERE movement_id=$1",
		id,
	)
	var movement model.Movement
	err := row.Scan(&movement.ID, &movement.Name)
	if err != nil {
		return nil, err
	}
	return &movement, nil
}

func (store MovementStore) List() ([]model.Movement, error) {
	rows, err := store.source.Query("SELECT movement_id, name FROM core.movements")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movements []model.Movement
	for rows.Next() {
		movement := model.Movement{}
		err := rows.Scan(&movement.ID, &movement.Name)
		if err != nil {
			return nil, err
		}
		movements = append(movements, movement)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return movements, nil
}

func (store MovementStore) Update(movement *model.Movement) error {
	_, err := store.source.Exec(
		"UPDATE core.movements SET name=$1 WHERE movement_id=$2",
		movement.Name,
		movement.ID,
	)
	return err
}

func (store MovementStore) Delete(id string) error {
	_, err := store.source.Exec(
		"DELETE FROM core.movements WHERE movement_id=$1",
		id,
	)
	return err
}
