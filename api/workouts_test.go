package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/boxtown/pupd/model"
	"github.com/boxtown/pupd/model/mock"
	"github.com/jmoiron/sqlx"
)

func TestListWorkouts(t *testing.T) {
	workouts := []*model.Workout{
		&model.Workout{
			ID:   "Test ID",
			Name: "Test Name",
		},
		&model.Workout{
			ID:   "Test ID 2",
			Name: "Test Name 2",
		},
	}
	store := mock.MockWorkoutStore{
		ListFn: func() ([]*model.Workout, error) {
			return workouts, nil
		},
	}
	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/workouts", nil)
	if err != nil {
		t.Fatal(err)
	}
	Router(nil, mock.MockStores{MockWorkoutStore: store}).ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Expected code %d, got %d", http.StatusOK, response.Code)
	}
	expected := bytes.Buffer{}
	err = json.NewEncoder(&expected).Encode(workouts)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(response.Body.Bytes(), expected.Bytes()) {
		t.Errorf("Response body does not match expected value")
	}
}

func TestListWorkoutsErrors(t *testing.T) {
	// Test store returns an error
	store := mock.MockWorkoutStore{
		ListFn: func() ([]*model.Workout, error) {
			return nil, errors.New("test")
		},
	}
	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/workouts", nil)
	if err != nil {
		t.Fatal(err)
	}
	Router(nil, mock.MockStores{MockWorkoutStore: store}).ServeHTTP(response, request)
	if response.Code != http.StatusInternalServerError {
		t.Errorf("Expected code %d, got %d", http.StatusInternalServerError, response.Code)
	}
}

func TestCreateWorkout(t *testing.T) {
	store := mock.MockWorkoutStore{
		CreateFn: func(workout *model.Workout) (string, error) {
			return "test", nil
		},
	}
	stores := mock.MockStores{MockWorkoutStore: store}
	dataSource := mock.MockDataSource{
		TransactionFn: func(handler func(sqlx.Ext) error) error {
			return handler(nil)
		},
	}

	response := httptest.NewRecorder()
	workout := model.Workout{Name: "Test Workout"}
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(workout)
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest("POST", "/workouts", &buf)
	if err != nil {
		t.Fatal(err)
	}
	Router(dataSource, stores).ServeHTTP(response, request)
	if response.Code != http.StatusCreated {
		t.Errorf("Expected code %d, got %d", http.StatusCreated, response.Code)
	}
	location := response.Header().Get("Location")
	if location != "/workouts/test" {
		t.Errorf("Expected Location header %s, got %s", "/workouts/test", location)
	}
}
