package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/vec-search/controllers/gateways"
)

func GatewayRoutes(r chi.Router) {
	r.Route("/gateway", func(r chi.Router) {
		r.Get("/{symbol}", gateways.ResearchSymbol)
		r.Get("/{symbol}/time_series", gateways.ResearchSymbolTimeSeries)
		r.Get("/{symbol}/news", gateways.ResearchSymbolNews)
		r.Get("/{symbol}/fifty_two", gateways.ResearchSymbolFiftyTwoWeek)
		r.Get(("/snapshot"), gateways.ResearchSymbolSnapshot)
	})
}
