// Code generated. This is hack to turn off linter for mocks. DO NOT EDIT.
// This will hopefully actually be code generated once github.com/golang/mock
// gets their shit together and merges in #28

package mock

import (
	"github.com/boxtown/pupd/model"
	"github.com/jmoiron/sqlx"
)

type MockStores struct {
	MockMovementStore
	MockWorkoutStore
	MockExerciseStore
}

func (stores MockStores) Movements(source sqlx.Ext) model.MovementStore {
	return stores.MockMovementStore
}

func (stores MockStores) Workouts(source sqlx.Ext) model.WorkoutStore {
	return stores.MockWorkoutStore
}

func (stores MockStores) Exercises(source sqlx.Ext) model.ExerciseStore {
	return stores.MockExerciseStore
}

type MockMovementStore struct {
	CreateFn    func(movement *model.Movement) (string, error)
	GetFn       func(id string) (*model.Movement, error)
	GetByNameFn func(name string) (*model.Movement, error)
	ListFn      func() ([]*model.Movement, error)
	UpdateFn    func(movement *model.Movement) error
	DeleteFn    func(id string) error
}

func (store MockMovementStore) Create(movement *model.Movement) (string, error) {
	return store.CreateFn(movement)
}

func (store MockMovementStore) Get(id string) (*model.Movement, error) {
	return store.GetFn(id)
}

func (store MockMovementStore) GetByName(name string) (*model.Movement, error) {
	return store.GetByNameFn(name)
}

func (store MockMovementStore) List() ([]*model.Movement, error) {
	return store.ListFn()
}

func (store MockMovementStore) Update(movement *model.Movement) error {
	return store.UpdateFn(movement)
}

func (store MockMovementStore) Delete(id string) error {
	return store.DeleteFn(id)
}

type MockWorkoutStore struct {
	CreateFn func(workout *model.Workout) (string, error)
	GetFn    func(id string) (*model.Workout, error)
	ListFn   func() ([]*model.Workout, error)
}

func (store MockWorkoutStore) Create(workout *model.Workout) (string, error) {
	return store.CreateFn(workout)
}

func (store MockWorkoutStore) Get(id string) (*model.Workout, error) {
	return store.GetFn(id)
}

func (store MockWorkoutStore) List() ([]*model.Workout, error) {
	return store.ListFn()
}

type MockExerciseStore struct {
	GetFn            func(id string) (*model.Exercise, error)
	GetByWorkoutIDFn func(id string) ([]*model.Exercise, error)
}

func (store MockExerciseStore) Get(id string) (*model.Exercise, error) {
	return store.GetFn(id)
}

func (store MockExerciseStore) GetByWorkoutID(id string) ([]*model.Exercise, error) {
	return store.GetByWorkoutIDFn(id)
}
