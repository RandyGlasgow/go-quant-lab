package controller

import (
	"encoding/json"
	"fmt"
	"github.com/vec-search/lib/market_stack"
	"github.com/vec-search/lib/market_stack/ms_utils"
	"github.com/vec-search/lib/redis"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/vec-search/lib/http/http_utils"
	"github.com/vec-search/lib/market_stack/ms_service"
)

func EndOfDayRoutes(r chi.Router) {

	r.Route("/eod", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

			rclient := redis.Redis.Client

			symbols, err := ms_utils.ExtractSymbolsFromQuery(r)

			if err != nil {
				http_utils.HttpCustomError(w, err)
			}

			if rclient.Exists(r.Context(), symbols).Val() == 1 {
				fmt.Println("symbol exists in cache")
				val, err := rclient.Get(r.Context(), symbols).Result()
				if err != nil {
					http_utils.HttpCustomError(w, err)
					return
				}
				// parse the data from the cache
				var jsonData = market_stack.EndOfDayService{}
				err = json.Unmarshal([]byte(val), &jsonData)

				if err != nil {
					http_utils.HttpCustomError(w, err)
					return
				}
				render.JSON(w, r, jsonData)
				return
			}

			resp, err := ms_service.MS_GetSymbolsEndOfDayByDate(r)

			if err != nil {
				http_utils.HttpCustomError(w, err)
			}

			var jsonData = market_stack.EndOfDayService{}
			err = http_utils.ExtractJsonFromResponseBody(resp, &jsonData)
			if err != nil {
				http_utils.HttpCustomError(w, err)
				return
			}

			binaryJson, err := json.Marshal(jsonData)
			if err != nil {
				http_utils.HttpCustomError(w, err)
				return
			}

			err = rclient.Set(r.Context(), symbols, binaryJson, 0).Err()

			render.JSON(w, r, jsonData)
			return
		})

		r.Get("/latest", func(w http.ResponseWriter, r *http.Request) {

			rclient := redis.Redis.Client

			symbols, err := ms_utils.ExtractSymbolsFromQuery(r)

			if err != nil {
				http_utils.HttpCustomError(w, err)
				return
			}

			if rclient.Exists(r.Context(), symbols).Val() == 1 {
				fmt.Println("\n\n\nsymbol exists in cache")

				val, err := rclient.Get(r.Context(), symbols).Result()
				if err != nil {
					http_utils.HttpCustomError(w, err)
					return
				}
				// parse the data from the cache
				var jsonData = market_stack.EndOfDayService{}
				err = json.Unmarshal([]byte(val), &jsonData)

				if err != nil {
					http_utils.HttpCustomError(w, err)
					return
				}
				render.JSON(w, r, jsonData)
				return
			}

			resp, err := ms_service.MS_GetSymbolsEndOfDayLatest(r)

			if err != nil {
				http_utils.HttpCustomError(w, err)
				return
			}

			var jsonData = market_stack.EndOfDayService{}
			err = http_utils.ExtractJsonFromResponseBody(resp, &jsonData)
			if err != nil {
				http_utils.HttpCustomError(w, err)
				return
			}

			binaryJson, err := json.Marshal(jsonData)
			if err != nil {
				http_utils.HttpCustomError(w, err)
				return
			}

			// set the exparatip to be tomorrow at 6am
			// this is to ensure that the data is always fresh
			// and that the cache is not stale
			tomorrow6am := time.Now().Add(24 * time.Hour).Truncate(24 * time.Hour).Add(6 * time.Hour)
			err = rclient.Set(r.Context(), symbols, binaryJson, tomorrow6am.Sub(time.Now())).Err()

			render.JSON(w, r, jsonData)
			return
		})
	})
}
