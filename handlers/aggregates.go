package handlers

import (
	"errors"
	"math"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/polygon-io/client-go/rest/iter"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/vec-search/lib"
	"github.com/vec-search/lib/http/http_utils"
)

func GetSymbolAggs(params *models.ListAggsParams, r *http.Request) *iter.Iter[models.Agg] {
	c := lib.PolyIoClient.Client
	res := c.ListAggs(r.Context(), &models.ListAggsParams{
		Ticker:     params.Ticker,
		Multiplier: params.Multiplier,
		Timespan:   params.Timespan,
		From:       params.From,
		To:         params.To,
	})

	return res
}

func FiftyTwoWeekHighLow(w http.ResponseWriter, r *http.Request) {
	ticker := chi.URLParam(r, "symbol")
	if ticker == "" {
		http_utils.HttpCustomError(w, errors.New("ticker is required"))
		return
	}

	today := time.Now()

	res := GetSymbolAggs(&models.ListAggsParams{
		Ticker:     ticker,
		Multiplier: 1,
		Timespan:   "week",
		From:       models.Millis(today.Add(-time.Hour * 24 * 365)),
		To:         models.Millis(today),
	}, r)

	data := struct {
		Ticker string  `json:"ticker"`
		High52 float64 `json:"high_52"`
		Low52  float64 `json:"low_52"`
		Avg52  float64 `json:"avg_52"`
	}{}

	var high = 0.0
	var low = math.MaxFloat64
	var avg = 0.0
	var count = 0

	for res.Next() {
		agg := res.Item()
		count++

		if agg.High > high {
			high = agg.High
		}
		if agg.Close < low {
			low = agg.Close
		}
		avg += agg.Close
	}

	data.Ticker = ticker
	data.High52 = high
	data.Low52 = low
	data.Avg52 = avg / float64(count)

	render.JSON(w, r, data)
}
