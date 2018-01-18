package model

import "errors"

// ErrBadIndices is returned by Model methods
// when supplied index parameters are invalid
var ErrBadIndices = errors.New("One or more of the supplied indices were out of bounds")

// Movement holds information relevant to any sort of
// exercise movement (i.e. Back Squat, 400m Run)
type Movement struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Unit is any unit of measurement (kg, s, m, etc.)
type Unit struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ExerciseSet is a single set of an exercise
// with a given number of reps, intensity range, unit of work
// and position within the set
type ExerciseSet struct {
	Pos          int      `json:"pos"`
	Reps         int      `json:"reps"`
	MinIntensity float32  `json:"minIntensity"`
	MaxIntensity *float32 `json:"maxIntensity,omitempty"`
	Unit         *Unit    `json:"unit"`
}

// Exercise is a collection of variable sets for a single movement
// (i.e. 2 x 5, 1 x 3 Back Squat)
type Exercise struct {
	ID       string         `json:"id"`
	Pos      int            `json:"pos"`
	Movement *Movement      `json:"movement"`
	Sets     []*ExerciseSet `json:"sets"`
}

// Workout is a named collection of exercises
type Workout struct {
	ID        string      `json:"id"`
	Name      string      `json:"name"`
	Exercises []*Exercise `json:"exercises"`
}
