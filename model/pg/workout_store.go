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

// Get attempts to retrieve a Workout from storage by
// its ID
func (store WorkoutStore) Get(id string) (*model.Workout, error) {
	row := store.source.QueryRowx(
		"SELECT name FROM core.workouts WHERE workout_id=$1",
		id,
	)
	workout := model.Workout{ID: id}
	if err := row.Scan(&workout.Name); err != nil {
		return nil, err
	}
	workoutExercises, err := store.exercises.GetByWorkoutID(id)
	if err != nil {
		return nil, err
	}
	workout.Exercises = workoutExercises
	return &workout, nil
}

// List lists all Workouts from storage
func (store WorkoutStore) List() ([]*model.Workout, error) {
	rows, err := store.source.Query("SELECT workout_id, name FROM core.workouts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workouts := []*model.Workout{}
	for rows.Next() {
		workout := model.Workout{}
		if err := rows.Scan(&workout.ID, &workout.Name); err != nil {
			return nil, err
		}
		workouts = append(workouts, &workout)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	for _, workout := range workouts {
		workoutExercises, err := store.exercises.GetByWorkoutID(workout.ID)
		if err != nil {
			return nil, err
		}
		workout.Exercises = workoutExercises
	}
	return workouts, nil
}
