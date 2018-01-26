package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/boxtown/pupd/model"
)

func listMovementsFn(store model.MovementStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movements, err := store.List()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(movements); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func createMovementFn(store model.MovementStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movement := model.Movement{}
		if err := json.NewDecoder(r.Body).Decode(&movement); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := store.Create(&movement)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Add("Location", fmt.Sprintf("/movements/%s", id))
		return
	}
}
