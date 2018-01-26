package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestListWorkouts(t *testing.T) {
	workouts := []*model.Workout{
		&model.Workout{
			ID:   "Test ID",
			Name: "Test Name",
			Exercises: []*model.Exercise{
				&model.Exercise{
					ID: "Test Exercise ID",
					Movement: &model.Movement{
						ID:   "Test Movement ID",
						Name: "Test Movement Name",
					},
					Sets: []*model.ExerciseSet{
						&model.ExerciseSet{
							Pos:  0,
							Reps: 5,
							Unit: &model.Unit{
								ID:   "Test Unit ID",
								Name: "Test Unit Name",
							},
						},
						&model.ExerciseSet{
							Pos:  1,
							Reps: 5,
							Unit: &model.Unit{
								ID:   "Test Unit ID",
								Name: "Test Unit Name",
							},
						},
					},
				},
			},
		},
		&model.Workout{
			ID:        "Test ID 2",
			Name:      "Test Name 2",
			Exercises: []*model.Exercise{},
		},
	}
	store := mockWorkoutStore{
		list: func() ([]*model.Workout, error) {
			return workouts, nil
		},
	}
	response := httptest.NewRecorder()
	listWorkoutsFn(store)(response, nil)
	if response.Code != http.StatusOK {
		t.Errorf("Expected code %d, got %d", http.StatusOK, response.Code)
	}
	expected := bytes.Buffer{}
	err := json.NewEncoder(&expected).Encode(workouts)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(response.Body.Bytes(), expected.Bytes()) {
		t.Errorf("Response body does not match expected value")
	}
}

func TestListWorkoutsErrors(t *testing.T) {
	// Test store returns an error
	store := mockWorkoutStore{
		list: func() ([]*model.Workout, error) {
			return nil, errors.New("test")
		},
	}
	response := httptest.NewRecorder()
	listWorkoutsFn(store)(response, nil)
	if response.Code != http.StatusInternalServerError {
		t.Errorf("Expected code %d, got %d", http.StatusInternalServerError, response.Code)
	}

	// Test store returns an un-encodable result
	store = mockWorkoutStore{
		list: func() ([]*model.Workout, error) {
			return nil, nil
		},
	}
	listWorkoutsFn(store)(response, nil)
	if response.Code != http.StatusInternalServerError {
		t.Errorf("Expected code %d, got %d", http.StatusInternalServerError, response.Code)
	}
}
