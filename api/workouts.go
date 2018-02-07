package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/boxtown/pupd/model"
	"github.com/jmoiron/sqlx"
)

func listWorkoutsFn(source model.DataSource, stores model.Stores) http.HandlerFunc {
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

func createWorkoutsFn(source model.DataSource, stores model.Stores) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		workout := model.Workout{}
		if err := json.NewDecoder(r.Body).Decode(&workout); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err := source.Transaction(func(tx sqlx.Ext) error {
			store := stores.Workouts(tx)
			id, err := store.Create(&workout)
			if err != nil {
				return err
			}
			w.WriteHeader(http.StatusCreated)
			w.Header().Set("Location", fmt.Sprintf("/workouts/%s", id))
			return nil
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
