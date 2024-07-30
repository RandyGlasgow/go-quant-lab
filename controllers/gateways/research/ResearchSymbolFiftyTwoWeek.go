package research

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/vec-search/lib"
)

func ResearchSymbolFiftyTwoWeek(w http.ResponseWriter, r *http.Request) {

	symbol := chi.URLParam(r, "symbol")

	// make a request to get the 52 week high and low buy looking at the last year of data using the year timespan

	client := lib.PolygonClient
	params := models.ListAggsParams{
		Ticker:     symbol,
		Multiplier: 365,
		Timespan:   models.Day,
		From:       models.Millis(time.Now().AddDate(-1, 1, 0)),
		To:         models.Millis(time.Now()),
	}

	resp := client.ListAggs(r.Context(), &params)

	// the data will already be calculated for us, all we need to do is remap it to the response it will return [] entries so we need to grab the first available one.
	data := struct {
		Ticker string  `json:"ticker"`
		High52 float64 `json:"high_52"`
		Low52  float64 `json:"low_52"`
		Avg52  float64 `json:"avg_52"`
	}{}

	value := resp.Item()

	for {
		if !resp.Next() {
			break
		}
		value = resp.Item()

		// print as json for now

		data.Ticker = symbol
		data.High52 = value.High
		data.Low52 = value.Low
		data.Avg52 = value.Close

	}

	render.JSON(w, r, data)
}
