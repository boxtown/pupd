package model

import "errors"

// ErrBadIndices is returned by Model methods
// when supplied index parameters are invalid
var ErrBadIndices = errors.New("One or more of the supplied indices were out of bounds")

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

// // Sets returns a slice of ExerciseSets
// // for the Exercise where the index of each
// // Set indicates it's position in the exercise.
// // This is reflected in the `Pos` field for each Set
// func (e Exercise) Sets() []ExerciseSet {
// 	dst := make([]ExerciseSet, len(e.sets))
// 	copy(dst, e.sets)
// 	return dst
// }

// // Push pushes a copy of a the supplied Set onto
// // the end of the slice of ExerciseSets for the given Exercise.
// // The pushed copy's position will be updated to reflect its
// // position within the slice
// func (e *Exercise) Push(set ExerciseSet) {
// 	set.Pos = len(e.sets)
// 	e.sets = append(e.sets, set)
// }

// // Swap swaps the position of two sets within the
// // Exercise. Returns an error if one or more of the indices
// // are invalid
// func (e *Exercise) Swap(i1, i2 int) error {
// 	if i1 < 0 || i2 < 0 || i1 >= len(e.sets) || i2 >= len(e.sets) {
// 		return ErrBadIndices
// 	}
// 	e.sets[i1].Pos = i2
// 	e.sets[i2].Pos = i1
// 	e.sets[i1], e.sets[i2] = e.sets[i2], e.sets[i1]
// 	return nil
// }

// Workout is a named collection of exercises
type Workout struct {
	ID        string
	Name      string
	Exercises []Exercise
}
