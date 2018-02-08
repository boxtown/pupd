package pg

import (
	"github.com/boxtown/pupd/model"
	"github.com/jmoiron/sqlx"
)

// ExerciseStore implements model.ExerciseStore
// using PostgreSQL
type ExerciseStore struct {
	*AbstractStore
	source    sqlx.Ext
	movements model.MovementStore
}

// NewExerciseStore returns a PostgreSQL-backed implementation
// of model.ExerciseStore
func NewExerciseStore(source sqlx.Ext, configs ...StoreConfig) model.ExerciseStore {
	store := &ExerciseStore{
		AbstractStore: &AbstractStore{
			idGen: UUIDV4Generator{},
		},
		source:    source,
		movements: NewMovementStore(source),
	}
	for _, config := range configs {
		config(store.AbstractStore)
	}
	return store
}

// Get attempts to retrieve an Exercise from storage by
// its ID
func (store ExerciseStore) Get(id string) (*model.Exercise, error) {
	row := store.source.QueryRowx(
		`SELECT e.pos, e.movement_id, m.name
			FROM core.exercises AS e
			INNER JOIN core.movements AS m ON e.movement_id=m.movement_id
			WHERE e.exercise_id=$1`,
		id,
	)
	exercise := model.Exercise{ID: id, Movement: &model.Movement{}}
	if err := row.Scan(
		&exercise.Pos,
		&exercise.Movement.ID,
		&exercise.Movement.Name,
	); err != nil {
		return nil, err
	}
	exerciseSets, err := store.getExerciseSets(id)
	if err != nil {
		return nil, err
	}
	exercise.Sets = exerciseSets
	return &exercise, nil
}

// GetByWorkoutID attempts to retrieves exercises for a given Workout
// by the Workout's ID
func (store ExerciseStore) GetByWorkoutID(id string) ([]*model.Exercise, error) {
	var exercises []*model.Exercise
	rows, err := store.source.Queryx(
		`SELECT e.exercise_id, e.pos, e.movement_id, m.name
			FROM core.exercises AS e
			INNER JOIN core.movements AS m ON e.movement_id=m.movement_id
			WHERE e.workout_id=$1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		exercise := model.Exercise{Movement: &model.Movement{}}
		if err := rows.Scan(
			&exercise.ID,
			&exercise.Pos,
			&exercise.Movement.ID,
			&exercise.Movement.Name,
		); err != nil {
			return nil, err
		}
		exercises = append(exercises, &exercise)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for _, exercise := range exercises {
		exerciseSets, err := store.getExerciseSets(exercise.ID)
		if err != nil {
			return nil, err
		}
		exercise.Sets = exerciseSets
	}
	return exercises, nil
}

func (store ExerciseStore) getExerciseSets(id string) ([]*model.ExerciseSet, error) {
	var exerciseSets []*model.ExerciseSet
	rows, err := store.source.Queryx(
		"SELECT e.pos, e.reps FROM core.exercise_sets AS e WHERE e.exercise_id=$1",
		id,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		set := model.ExerciseSet{}
		if err := rows.Scan(
			&set.Pos,
			&set.Reps,
		); err != nil {
			return nil, err
		}
		exerciseSets = append(exerciseSets, &set)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return exerciseSets, nil
}
