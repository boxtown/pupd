package api

import (
	"encoding/json"
	"net/http"

	"github.com/boxtown/pupd/model"
	"github.com/boxtown/pupd/model/pg"
)

func listWorkoutsFn(source *pg.DataSource, stores model.Stores) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store := stores.Workouts(source)
		workouts, err := store.List()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(workouts)
	}
}

func createWorkoutsFn(source *pg.DataSource, stores model.Stores) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "/workouts/test")
		w.WriteHeader(http.StatusCreated)
	}
}
