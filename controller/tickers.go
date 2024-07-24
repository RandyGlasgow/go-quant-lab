package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/vec-search/lib/http/http_utils"
	"github.com/vec-search/lib/market_stack"
	"github.com/vec-search/lib/market_stack/ms_service"
	"net/http"
)

func TickerRoutes(r chi.Router) {
	r.Route("/tickers/{symbol}", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			resp, err := ms_service.MS_GetSymbolInfo(r)

			if err != nil {
				http_utils.HttpCustomError(w, err)
			}

			var json = market_stack.TickerInfo{}
			err = http_utils.ExtractJsonFromResponseBody(resp, &json)
			if err != nil {
				http_utils.HttpCustomError(w, err)
				return
			}
			render.JSON(w, r, json)
		})
	})
}
