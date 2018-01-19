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

type mockMovementStore struct {
	create    func(movement *model.Movement) (string, error)
	get       func(id string) (*model.Movement, error)
	getByName func(name string) (*model.Movement, error)
	list      func() ([]model.Movement, error)
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

func (store mockMovementStore) List() ([]model.Movement, error) {
	return store.list()
}

func (store mockMovementStore) Update(movement *model.Movement) error {
	return store.update(movement)
}

func (store mockMovementStore) Delete(id string) error {
	return store.delete(id)
}

func TestListMovements(t *testing.T) {
	movements := []model.Movement{
		model.Movement{ID: "Test ID", Name: "Test Name"},
		model.Movement{ID: "Test ID 2", Name: "Test Name 2"},
	}
	store := mockMovementStore{
		list: func() ([]model.Movement, error) {
			return movements, nil
		},
	}
	response := httptest.NewRecorder()
	listMovementsFn(store)(response, nil)
	if response.Code != http.StatusOK {
		t.Errorf("Expected code %d, got %d", http.StatusOK, response.Code)
	}
	var expected bytes.Buffer
	json.NewEncoder(&expected).Encode(movements)
	if !bytes.Equal(response.Body.Bytes(), expected.Bytes()) {
		t.Errorf("Response body does not match expected value")
	}
}

func TestListMovementErrors(t *testing.T) {
	// Test store throws an error
	store := mockMovementStore{
		list: func() ([]model.Movement, error) {
			return nil, errors.New("test")
		},
	}
	response := httptest.NewRecorder()
	listMovementsFn(store)(response, nil)
	if response.Code != http.StatusInternalServerError {
		t.Errorf("Expedcted code %d, got %d", http.StatusInternalServerError, response.Code)
	}

	// Test store returns un-encodable result
	store = mockMovementStore{
		list: func() ([]model.Movement, error) {
			return nil, nil
		},
	}
	listMovementsFn(store)(response, nil)
	if response.Code != http.StatusInternalServerError {
		t.Errorf("Expedcted code %d, got %d", http.StatusInternalServerError, response.Code)
	}
}
