package controller

import (
	"github.com/vec-search/lib/market_stack"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/vec-search/lib/http/http_utils"
	"github.com/vec-search/lib/market_stack/ms_service"
)

func EndOfDayRoutes(r chi.Router) {

	r.Route("/eod", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			resp, err := ms_service.MS_GetSymbolsEndOfDayByDate(r)

			if err != nil {
				http_utils.HttpCustomError(w, err)
			}

			var json = market_stack.EndOfDayService{}
			err = http_utils.ExtractJsonFromResponseBody(resp, &json)
			if err != nil {
				http_utils.HttpCustomError(w, err)
				return
			}
			render.JSON(w, r, json)
		})

		r.Get("/latest", func(w http.ResponseWriter, r *http.Request) {
			resp, err := ms_service.MS_GetSymbolsEndOfDayLatest(r)

			if err != nil {
				http_utils.HttpCustomError(w, err)
			}

			var json = market_stack.EndOfDayService{}
			err = http_utils.ExtractJsonFromResponseBody(resp, &json)
			if err != nil {
				http_utils.HttpCustomError(w, err)
				return
			}
			render.JSON(w, r, json)
		})
	})
}
