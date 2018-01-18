package pg

import (
	"github.com/boxtown/pupd/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// ExerciseStore implements model.ExerciseStore
// using PostgreSQL
type ExerciseStore struct {
	source    sqlx.Ext
	movements model.MovementStore
	units     model.UnitStore
}

// NewExerciseStore returns a PostgreSQL-backed implementation
// of model.ExerciseStore
func NewExerciseStore(source sqlx.Ext) model.ExerciseStore {
	return &ExerciseStore{
		source:    source,
		movements: NewMovementStore(source),
		units:     NewUnitStore(source),
	}
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

	if err = store.insertExercise(id, workoutID, exercise); err != nil {
		return "", err
	}
	for _, set := range exercise.Sets {
		// TODO: ensure proper Pos field?
		if err = store.insertExerciseSet(id, set); err != nil {
			return "", err
		}
	}
	return id, nil
}

func (store ExerciseStore) insertExercise(id, workoutID string, exercise *model.Exercise) error {
	if len(exercise.Movement.ID) == 0 {
		movement, err := store.movements.GetByName(exercise.Movement.Name)
		if err != nil {
			return err
		}
		exercise.Movement = movement
	}
	_, err := store.source.Exec(
		"INSERT INTO core.exercises (exercise_id, workout_id, pos, movement_id) VALUES ($1, $2, $3, $4)",
		id,
		workoutID,
		exercise.Pos,
		exercise.Movement.ID,
	)
	return err
}

func (store ExerciseStore) insertExerciseSet(id string, set *model.ExerciseSet) error {
	// TODO: ensure proper Pos field?
	if len(set.Unit.ID) == 0 {
		unit, err := store.units.GetByName(set.Unit.Name)
		if err != nil {
			return err
		}
		set.Unit = unit
	}
	_, err := store.source.Exec(
		`INSERT INTO core.exercise_sets (exercise_id, pos, reps, min_intensity, max_intensity, unit_id)
			VALUES ($1, $2, $3, $4, $5, $6)`,
		id,
		set.Pos,
		set.Reps,
		set.MinIntensity,
		set.MaxIntensity,
		set.Unit.ID,
	)
	return err
}
