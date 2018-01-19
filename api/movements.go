package api

import (
	"encoding/json"
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
