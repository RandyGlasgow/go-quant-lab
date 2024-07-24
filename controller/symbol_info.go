package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/vec-search/handlers"
)

func SymbolInfoRoutes(r chi.Router) {
	r.Route("/symbol_info", func(r chi.Router) {
		r.Get("/{date}", handlers.SymbolInfo)
	})
}
