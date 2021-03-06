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

// Workouts returns a PostgreSQL backed `model.WorkoutStore`
func (store Stores) Workouts(source sqlx.Ext) model.WorkoutStore {
	return NewWorkoutStore(source)
}

// Exercises returns a PostgreSQL backed `model.ExerciseStore`
func (store Stores) Exercises(source sqlx.Ext) model.ExerciseStore {
	return NewExerciseStore(source)
}

// AbstractStore represents the base fields all stores
// must have
type AbstractStore struct {
	idGen IDGenerator
}

// StoreConfig represents a configuration function for
// inheritors of AbstractStore
type StoreConfig func(*AbstractStore)

// WithIDGenerator generates a StoreConfig that
// sets the IDGenerator for a store to the supplied
// IDGenerator
func WithIDGenerator(idGen IDGenerator) StoreConfig {
	return func(store *AbstractStore) {
		store.idGen = idGen
	}
}
