package controller

import (
	"github.com/vec-search/handlers"

	"github.com/go-chi/chi/v5"
)

func EndOfDayRoutes(r chi.Router) {
	r.Route("/eod", func(r chi.Router) {
		r.Get("/latest", handlers.EndOfDayLatest)
		r.Get("/{date}", handlers.EndOfDayByDate)
	})
}
