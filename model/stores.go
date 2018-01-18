package model

// MovementStore defines an interface for a store of Movements
type MovementStore interface {
	Create(movement *Movement) (string, error)
	Get(id string) (*Movement, error)
	GetByName(name string) (*Movement, error)
	List() ([]Movement, error)
	Update(movement *Movement) error
	Delete(id string) error
}

// UnitStore defines an interface for a store of Units
type UnitStore interface {
	Create(unit *Unit) (string, error)
	Get(id string) (*Unit, error)
	GetByName(name string) (*Unit, error)
	List() ([]Unit, error)
	Update(unit *Unit) error
	Delete(id string) error
}

// WorkoutStore defines an interface for a store of Workouts
type WorkoutStore interface {
	Create(workout *Workout) (string, error)
}

// ExerciseStore defines an interface for a store of Exercises
type ExerciseStore interface {
	Create(workoutID string, exercise *Exercise) (string, error)
	Get(id string) (*Exercise, error)
}
