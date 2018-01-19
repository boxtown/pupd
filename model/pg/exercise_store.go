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

func (store ExerciseStore) insertExercise(id, workoutID string, exercise *model.Exercise) error {
	movementID := exercise.Movement.ID
	if len(movementID) == 0 {
		movement, err := store.movements.GetByName(exercise.Movement.Name)
		if err != nil {
			return err
		}
		movementID = movement.ID
	}
	_, err := store.source.Exec(
		"INSERT INTO core.exercises (exercise_id, workout_id, pos, movement_id) VALUES ($1, $2, $3, $4)",
		id,
		workoutID,
		exercise.Pos,
		movementID,
	)
	return err
}

func (store ExerciseStore) insertExerciseSet(id string, set *model.ExerciseSet) error {
	// TODO: ensure proper Pos field?
	unitID := set.Unit.ID
	if len(unitID) == 0 {
		unit, err := store.units.GetByName(set.Unit.Name)
		if err != nil {
			return err
		}
		unitID = unit.ID
	}
	_, err := store.source.Exec(
		`INSERT INTO core.exercise_sets (exercise_id, pos, reps, min_intensity, max_intensity, unit_id)
			VALUES ($1, $2, $3, $4, $5, $6)`,
		id,
		set.Pos,
		set.Reps,
		set.MinIntensity,
		set.MaxIntensity,
		unitID,
	)
	return err
}

func (store ExerciseStore) getExerciseSets(id string) ([]*model.ExerciseSet, error) {
	var exerciseSets []*model.ExerciseSet
	rows, err := store.source.Queryx(
		`SELECT e.pos, e.reps, e.min_intensity, e.max_intensity, e.unit_id, u.name
			FROM core.exercise_sets AS e
			INNER JOIN core.units AS u ON e.unit_id=u.unit_id
			WHERE e.exercise_id=$1`,
		id,
	)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		set := model.ExerciseSet{Unit: &model.Unit{}}
		if err := rows.Scan(
			&set.Pos,
			&set.Reps,
			&set.MinIntensity,
			&set.MaxIntensity,
			&set.Unit.ID,
			&set.Unit.Name,
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
