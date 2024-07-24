package controller

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/vec-search/lib/http/http_utils"
	"github.com/vec-search/lib/market_stack"
	"github.com/vec-search/lib/market_stack/ms_service"
	"github.com/vec-search/lib/redis"
	"net/http"
)

func TickerRoutes(r chi.Router) {
	r.Route("/tickers/{symbol}", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

			rclient := redis.Redis.Client
			// check that the symbol is in the cache
			symbol := chi.URLParam(r, "symbol")
			fmt.Println("symbol:", symbol)

			// check if the symbol is in the cache
			if rclient.Exists(r.Context(), symbol).Val() == 1 {
				fmt.Println("symbol exists in cache")
				val, err := rclient.Get(r.Context(), symbol).Result()
				if err != nil {
					http_utils.HttpCustomError(w, err)
					return
				}
				// parse the data from the cache

				var jsonData = market_stack.TickerInfo{}
				err = json.Unmarshal([]byte(val), &jsonData)
				if err != nil {
					http_utils.HttpCustomError(w, err)
					return
				}
				render.JSON(w, r, jsonData)
				return
			}

			// request the data from the marketstack api
			resp, err := ms_service.MS_GetSymbolInfo(r)

			if err != nil {
				http_utils.HttpCustomError(w, err)
				return
			}

			// read the response body
			// redis: can't marshal market_stack.TickerInfo (implement encoding.BinaryMarshaler)
			var jsonData = market_stack.TickerInfo{}
			err = http_utils.ExtractJsonFromResponseBody(resp, &jsonData)

			if err != nil {
				http_utils.HttpCustomError(w, err)
				return
			}
			binaryJsonData, err := json.Marshal(jsonData)

			// store the data in the cache
			err = rclient.Set(r.Context(), symbol, binaryJsonData, 0).Err()
			if err != nil {
				http_utils.HttpCustomError(w, err)
				return
			}

			render.JSON(w, r, jsonData)
			return
		})
	})
}
