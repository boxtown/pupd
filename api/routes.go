package api

import (
	"net/http"

	"github.com/boxtown/pupd/model/pg"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
)

// Router returns a router http.Handler
// to be served by `http.ListenAndServe`
func Router(source sqlx.Ext) http.Handler {
	r := chi.NewRouter()
	r.Route("/movements", func(r chi.Router) {
		r.Get("/", listMovementsFn(pg.NewMovementStore(source)))
	})
	return r
}
