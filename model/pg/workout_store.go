package pg

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/boxtown/pupd/model"
	"github.com/jmoiron/sqlx"
)

// WorkoutStore implements model.WorkoutStore
// using PostgreSQL
type WorkoutStore struct {
	*AbstractStore
	source    sqlx.Ext
	exercises model.ExerciseStore
}

// NewWorkoutStore returns a PostgreSQL-backed implementation
// of model.WorkoutStore
func NewWorkoutStore(source sqlx.Ext, configs ...StoreConfig) model.WorkoutStore {
	store := &WorkoutStore{
		AbstractStore: &AbstractStore{
			idGen: UUIDV4Generator{},
		},
		source:    source,
		exercises: NewExerciseStore(source),
	}
	for _, config := range configs {
		config(store.AbstractStore)
	}
	return store
}

// Create attempts to create a record for the given Workout
// in the store. Any related records (i.e. for Exercises and Sets)
// will also be created. A v4 UUID will be assigned to the Workout
// and is returned by this method.
func (store WorkoutStore) Create(workout *model.Workout) (string, error) {
	id, err := store.idGen.Generate()
	if err != nil {
		return "", err
	}

	if _, err = store.source.Exec(
		"INSERT INTO core.workouts (workout_id, name) VALUES ($1, $2)",
		id,
		workout.Name,
	); err != nil {
		return "", err
	}
	if err = store.createExercises(id, workout.Exercises); err != nil {
		return "", err
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
	return workouts, nil
}

func (store WorkoutStore) createExercises(workoutID string, exercises []*model.Exercise) error {
	if len(exercises) == 0 {
		return nil
	}
	exerciseIDs := make([]string, len(exercises))
	for i := range exercises {
		id, err := store.idGen.Generate()
		if err != nil {
			return err
		}
		exerciseIDs[i] = id
	}
	query, args := buildCreateExercisesQuery(workoutID, exerciseIDs, exercises)
	if _, err := store.source.Exec(query, args...); err != nil {
		return err
	}
	query, args = buildCreateExerciseSetsQuery(exerciseIDs, exercises)
	if _, err := store.source.Exec(query, args...); err != nil {
		return err
	}
	return nil
}

func buildCreateExercisesQuery(workoutID string, exerciseIDs []string, exercises []*model.Exercise) (string, []interface{}) {
	query := bytes.Buffer{}
	query.WriteString("INSERT INTO core.exercises (exercise_id, workout_id, pos, movement_id) VALUES")
	args := []interface{}{}

	values := make([]string, len(exercises))
	for i, exercise := range exercises {
		values[i] = fmt.Sprintf(" ($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4)
		args = append(args, exerciseIDs[i], workoutID, i, exercise.Movement.ID)
	}
	query.WriteString(strings.Join(values, ","))
	return query.String(), args
}

func buildCreateExerciseSetsQuery(exerciseIDs []string, exercises []*model.Exercise) (string, []interface{}) {
	query := bytes.Buffer{}
	query.WriteString("INSERT INTO core.exercises_sets (exercise_id, pos, reps) VALUES")
	args := []interface{}{}

	values := []string{}
	k := 0
	for i, exercise := range exercises {
		for j, set := range exercise.Sets {
			values = append(values, fmt.Sprintf(" ($%d, $%d, $%d)", k*3+1, k*3+2, k*3+3))
			args = append(args, exerciseIDs[i], j, set.Reps)
			k++
		}
	}
	query.WriteString(strings.Join(values, ","))
	return query.String(), args
}
