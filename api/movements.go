package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/boxtown/pupd/model"
	"github.com/boxtown/pupd/model/pg"
)

func listMovementsFn(source *pg.DataSource, stores model.Stores) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		store := stores.Movements(source)
		movements, err := store.List()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(movements)
	}
}

func createMovementFn(source *pg.DataSource, stores model.Stores) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		movement := model.Movement{}
		if err := json.NewDecoder(r.Body).Decode(&movement); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		store := stores.Movements(source)
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
