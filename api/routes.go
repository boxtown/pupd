package api

import (
	"net/http"

	"github.com/boxtown/pupd/model/pg"

	"github.com/boxtown/pupd/model"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

// Router returns a router http.Handler
// to be served by `http.ListenAndServe`
func Router(source *pg.DataSource, stores model.Stores) http.Handler {
	r := chi.NewRouter()

	cors := cors.Default()
	r.Use(cors.Handler)

	r.Route("/movements", func(r chi.Router) {
		r.Get("/", listMovementsFn(source, stores))
		r.Post("/", createMovementFn(source, stores))
	})
	r.Route("/workouts", func(r chi.Router) {
		r.Get("/", listWorkoutsFn(source, stores))
	})
	return r
}
