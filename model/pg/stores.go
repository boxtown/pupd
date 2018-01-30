package pg

import (
	"github.com/boxtown/pupd/model"
	"github.com/jmoiron/sqlx"
)

// Stores is the PostgreSQL backed `model.Stores`
// implementation
type Stores struct{}

// Movements returns a PostgreSQL backed `model.MovementStore`
func (store Stores) Movements(source sqlx.Ext) model.MovementStore {
	return NewMovementStore(source)
}

// Units returns a PostgreSQL backed `model.UnitStore`
func (store Stores) Units(source sqlx.Ext) model.UnitStore {
	return NewUnitStore(source)
}

// Workouts returns a PostgreSQL backed `model.WorkoutStore`
func (store Stores) Workouts(source sqlx.Ext) model.WorkoutStore {
	return NewWorkoutStore(source)
}

// Exercises returns a PostgreSQL backed `model.ExerciseStore`
func (store Stores) Exercises(source sqlx.Ext) model.ExerciseStore {
	return NewExerciseStore(source)
}
