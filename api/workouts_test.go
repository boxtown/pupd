package api

import (
	"github.com/boxtown/pupd/model"
)

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
