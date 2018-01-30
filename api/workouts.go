package api

import (
	"encoding/json"
	"net/http"

	"github.com/boxtown/pupd/model"
)

func listWorkoutsFn(store model.WorkoutStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		workouts, err := store.List()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(workouts); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
