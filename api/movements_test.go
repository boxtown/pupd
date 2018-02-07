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
)

func TestListMovements(t *testing.T) {
	movements := []*model.Movement{
		&model.Movement{ID: "Test ID", Name: "Test Name"},
		&model.Movement{ID: "Test ID 2", Name: "Test Name 2"},
	}
	store := mock.MockMovementStore{
		ListFn: func() ([]*model.Movement, error) {
			return movements, nil
		},
	}
	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/movements", nil)
	if err != nil {
		t.Fatal(err)
	}
	Router(nil, mock.MockStores{MockMovementStore: store}).ServeHTTP(response, request)
	if response.Code != http.StatusOK {
		t.Errorf("Expected code %d, got %d", http.StatusOK, response.Code)
	}
	expected := bytes.Buffer{}
	err = json.NewEncoder(&expected).Encode(movements)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(response.Body.Bytes(), expected.Bytes()) {
		t.Errorf("Response body does not match expected value")
	}
}

func TestListMovementsErrors(t *testing.T) {
	// Test store returns an error
	store := mock.MockMovementStore{
		ListFn: func() ([]*model.Movement, error) {
			return nil, errors.New("test")
		},
	}
	response := httptest.NewRecorder()
	request, err := http.NewRequest("GET", "/movements", nil)
	if err != nil {
		t.Fatal(err)
	}
	Router(nil, mock.MockStores{MockMovementStore: store}).ServeHTTP(response, request)
	if response.Code != http.StatusInternalServerError {
		t.Errorf("Expected code %d, got %d", http.StatusInternalServerError, response.Code)
	}
}

func TestCreateMovement(t *testing.T) {
	store := mock.MockMovementStore{
		CreateFn: func(*model.Movement) (string, error) {
			return "test", nil
		},
	}
	response := httptest.NewRecorder()
	movement := model.Movement{Name: "Test Movement"}
	buf := bytes.Buffer{}
	err := json.NewEncoder(&buf).Encode(movement)
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest("POST", "/movements", &buf)
	if err != nil {
		t.Fatal(err)
	}
	Router(nil, mock.MockStores{MockMovementStore: store}).ServeHTTP(response, request)
	if response.Code != http.StatusCreated {
		t.Errorf("Expected code %d, got %d", http.StatusCreated, response.Code)
	}
	location := response.Header().Get("Location")
	if location != "/movements/test" {
		t.Errorf("Expected Location header %s, got %s", "/movements/test", location)
	}
}

func TestCreateMovementErrors(t *testing.T) {
	// Test bad JSON
	store := mock.MockMovementStore{
		CreateFn: func(*model.Movement) (string, error) {
			return "test", nil
		},
	}
	response := httptest.NewRecorder()
	buf := bytes.Buffer{}
	_, err := buf.WriteString("Not JSON")
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest("POST", "/movements", &buf)
	if err != nil {
		t.Fatal(err)
	}
	Router(nil, mock.MockStores{MockMovementStore: store}).ServeHTTP(response, request)
	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected code %d, got %d", http.StatusBadRequest, response.Code)
	}

	// Test store returns an error
	store = mock.MockMovementStore{
		CreateFn: func(*model.Movement) (string, error) {
			return "", errors.New("test")
		},
	}
	response = httptest.NewRecorder()
	movement := model.Movement{Name: "Test Movement"}
	buf = bytes.Buffer{}
	err = json.NewEncoder(&buf).Encode(movement)
	if err != nil {
		t.Fatal(err)
	}
	request, err = http.NewRequest("POST", "/movements", &buf)
	if err != nil {
		t.Fatal(err)
	}
	Router(nil, mock.MockStores{MockMovementStore: store}).ServeHTTP(response, request)
	if response.Code != http.StatusInternalServerError {
		t.Errorf("Expected code %d, got %d", http.StatusInternalServerError, response.Code)
	}
}
