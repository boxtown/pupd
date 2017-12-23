package model

type Movement struct {
	ID   string
	Name string
}

type Unit struct {
	ID   string
	Name string
}

type ExerciseSet struct {
	Pos          int
	Reps         int
	MinIntensity float32
	MaxIntensity *float32
	Unit         Unit
}

type Exercise struct {
	ID       string
	Pos      int
	Movement Movement
	Sets     []ExerciseSet
}

type Workout struct {
	ID        string
	Name      string
	Exercises []Exercise
}
