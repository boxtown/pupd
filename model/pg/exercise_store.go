package pg

import (
	"github.com/boxtown/pupd/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ExerciseStore implements model.ExerciseStore
// using PostgreSQL
type ExerciseStore struct {
	source sqlx.Ext
}

// NewExerciseStore returns a PostgreSQL-backed implementation
// of model.ExerciseStore
func NewExerciseStore(source sqlx.Ext) model.ExerciseStore {
	return &ExerciseStore{source: source}
}

// Create attempts to create a record for the given Exercise
// in the store. Any related records (i.e. for Sets) will also
// be created. A v4 UUID will be assigned to the Exercise
// and is returned by this method.
func (store ExerciseStore) Create(workoutID string, exercise *model.Exercise) (string, error) {
	raw, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	id := raw.String()

	if _, err = store.source.Exec(
		"INSERT INTO core.exercises (exercise_id, workout_id, pos, movement_id) VALUES ($1, $2, $3, $4)",
		id,
		workoutID,
		exercise.Pos,
		exercise.Movement.ID,
	); err != nil {
		return "", err
	}

	// TODO: create sets
	return id, nil
}
