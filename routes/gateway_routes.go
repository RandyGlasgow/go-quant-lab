package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/vec-search/controllers/gateways/research"
)

func GatewayRoutes(r chi.Router) {

	r.Route("/gateway", func(r chi.Router) {
		r.Get("/{symbol}", research.ResearchSymbolTickerDetails)
		r.Get("/{symbol}/time_series", research.ResearchSymbolTimeSeries)
		r.Get("/{symbol}/news", research.ResearchSymbolNews)
		r.Get("/{symbol}/fifty_two", research.ResearchSymbolFiftyTwoWeek)
		r.Get("/{symbol}/financials", research.ResearchSymbolSnapshot)
		r.Get(("/snapshot"), research.ResearchSymbolSnapshot)
	})
}
