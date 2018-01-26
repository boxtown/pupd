package api

import (
	"net/http"

	"github.com/boxtown/pupd/model"
)

func listWorkoutsFn(store model.WorkoutStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
