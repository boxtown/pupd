package model

// Movement holds information relevant to any sort of
// exercise movement (i.e. Back Squat, 400m Run)
type Movement struct {
	ID   string
	Name string
}

// Unit is any unit of measurement (kg, s, m, etc.)
type Unit struct {
	ID   string
	Name string
}

// ExerciseSet is a single set of an exercise
// with a given number of reps, intensity range, unit of work
// and position within the set
type ExerciseSet struct {
	Pos          int
	Reps         int
	MinIntensity float32
	MaxIntensity *float32
	Unit         Unit
}

// Exercise is a collection of variable sets for a single movement
// (i.e. 2 x 5, 1 x 3 Back Squat)
type Exercise struct {
	ID       string
	Pos      int
	Movement Movement
	Sets     []ExerciseSet
}

// Workout is a named collection of exercises
type Workout struct {
	ID        string
	Name      string
	Exercises []Exercise
}
