package api

import (
	"github.com/boxtown/pupd/model"
	"github.com/jmoiron/sqlx"
)

type mockStores struct {
	mockMovementStore
	mockUnitStore
	mockWorkoutStore
	mockExerciseStore
}

func (stores mockStores) Movements(source sqlx.Ext) model.MovementStore {
	return stores.mockMovementStore
}

func (stores mockStores) Units(source sqlx.Ext) model.UnitStore {
	return stores.mockUnitStore
}

func (stores mockStores) Workouts(source sqlx.Ext) model.WorkoutStore {
	return stores.mockWorkoutStore
}

func (stores mockStores) Exercises(source sqlx.Ext) model.ExerciseStore {
	return stores.mockExerciseStore
}

type mockMovementStore struct {
	create    func(movement *model.Movement) (string, error)
	get       func(id string) (*model.Movement, error)
	getByName func(name string) (*model.Movement, error)
	list      func() ([]*model.Movement, error)
	update    func(movement *model.Movement) error
	delete    func(id string) error
}

func (store mockMovementStore) Create(movement *model.Movement) (string, error) {
	return store.create(movement)
}

func (store mockMovementStore) Get(id string) (*model.Movement, error) {
	return store.get(id)
}

func (store mockMovementStore) GetByName(name string) (*model.Movement, error) {
	return store.getByName(name)
}

func (store mockMovementStore) List() ([]*model.Movement, error) {
	return store.list()
}

func (store mockMovementStore) Update(movement *model.Movement) error {
	return store.update(movement)
}

func (store mockMovementStore) Delete(id string) error {
	return store.delete(id)
}

type mockUnitStore struct {
	create    func(unit *model.Unit) (string, error)
	get       func(id string) (*model.Unit, error)
	getByName func(name string) (*model.Unit, error)
	list      func() ([]*model.Unit, error)
	update    func(unit *model.Unit) error
	delete    func(id string) error
}

func (store mockUnitStore) Create(unit *model.Unit) (string, error) {
	return store.create(unit)
}

func (store mockUnitStore) Get(id string) (*model.Unit, error) {
	return store.get(id)
}

func (store mockUnitStore) GetByName(name string) (*model.Unit, error) {
	return store.getByName(name)
}

func (store mockUnitStore) List() ([]*model.Unit, error) {
	return store.list()
}

func (store mockUnitStore) Update(unit *model.Unit) error {
	return store.update(unit)
}

func (store mockUnitStore) Delete(id string) error {
	return store.delete(id)
}

type mockWorkoutStore struct {
	create func(workout *model.Workout) (string, error)
	get    func(id string) (*model.Workout, error)
	list   func() ([]*model.Workout, error)
}

func (store mockWorkoutStore) Create(workout *model.Workout) (string, error) {
	return store.create(workout)
}

func (store mockWorkoutStore) Get(id string) (*model.Workout, error) {
	return store.get(id)
}

func (store mockWorkoutStore) List() ([]*model.Workout, error) {
	return store.list()
}

type mockExerciseStore struct {
	create         func(workoutID string, exercise *model.Exercise) (string, error)
	get            func(id string) (*model.Exercise, error)
	getByWorkoutID func(id string) ([]*model.Exercise, error)
}

func (store mockExerciseStore) Create(workoutID string, exercise *model.Exercise) (string, error) {
	return store.create(workoutID, exercise)
}

func (store mockExerciseStore) Get(id string) (*model.Exercise, error) {
	return store.get(id)
}

func (store mockExerciseStore) GetByWorkoutID(id string) ([]*model.Exercise, error) {
	return store.getByWorkoutID(id)
}
