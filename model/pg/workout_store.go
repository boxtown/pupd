package pg

import (
	"github.com/boxtown/pupd/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// WorkoutStore implements model.WorkoutStore
// using PostgreSQL
type WorkoutStore struct {
	source    sqlx.Ext
	exercises model.ExerciseStore
}

// NewWorkoutStore returns a PostgreSQL-backed implementation
// of model.WorkoutStore
func NewWorkoutStore(source sqlx.Ext) model.WorkoutStore {
	return &WorkoutStore{
		source:    source,
		exercises: NewExerciseStore(source),
	}
}

// Create attempts to create a record for the given Workout
// in the store. Any related records (i.e. for Exercises and Sets)
// will also be created. A v4 UUID will be assigned to the Workout
// and is returned by this method.
func (store WorkoutStore) Create(workout *model.Workout) (string, error) {
	raw, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	id := raw.String()

	if _, err = store.source.Exec(
		"INSERT INTO core.workouts (workout_id, name) VALUES ($1, $2)",
		id,
		workout.Name,
	); err != nil {
		return "", err
	}
	for _, exercise := range workout.Exercises {
		// TODO: ensure proper Pos field?
		if _, err = store.exercises.Create(id, exercise); err != nil {
			return "", err
		}
	}
	return id, nil
}
