package model

import "github.com/jmoiron/sqlx"

// Stores is a factory class for retrieving
// stores given a data source
type Stores interface {
	Movements(source sqlx.Ext) MovementStore
	Workouts(source sqlx.Ext) WorkoutStore
	Exercises(source sqlx.Ext) ExerciseStore
}

// MovementStore defines an interface for a store of Movements
type MovementStore interface {
	Create(movement *Movement) (string, error)
	Get(id string) (*Movement, error)
	GetByName(name string) (*Movement, error)
	List() ([]*Movement, error)
	Update(movement *Movement) error
	Delete(id string) error
}

// WorkoutStore defines an interface for a store of Workouts
type WorkoutStore interface {
	Create(workout *Workout) (string, error)
	Get(id string) (*Workout, error)
	List() ([]*Workout, error)
}

// ExerciseStore defines an interface for a store of Exercises
type ExerciseStore interface {
	Get(id string) (*Exercise, error)
	GetByWorkoutID(id string) ([]*Exercise, error)
}
