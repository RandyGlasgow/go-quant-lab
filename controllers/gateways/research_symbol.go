package gateways

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/vec-search/lib"
	"github.com/vec-search/lib/http/http_utils"
)

/**
 * @api {get} /research_symbol Research Symbol
 * @apiName ResearchSymbol
 * @apiGroup Research
 * Returns the fundamental data for a symbol.
 */
func ResearchSymbol(w http.ResponseWriter, r *http.Request) {
	symbol := chi.URLParam(r, "symbol")

	client := lib.PolygonClient
	params := models.GetTickerDetailsParams{
		Ticker: symbol,
	}

	data, err := client.GetTickerDetails(r.Context(), &params)
	if err != nil {
		http_utils.HttpStandardError(w, http.StatusBadRequest)
	}

	// Get the ticker details
	render.JSON(w, r, data)
}

func ResearchSymbolTimeSeries(w http.ResponseWriter, r *http.Request) {
	symbol := chi.URLParam(r, "symbol")
	// need the measure
	measureQuery := r.URL.Query().Get("measure")
	// need the multiplier to determine the timespan of the measure
	multiplier, err := strconv.Atoi(r.URL.Query().Get("multiplier"))
	if err != nil {
		http_utils.HttpCustomError(w, errors.New("invalid multiplier"))
		return
	}

	delta, err := strconv.Atoi(r.URL.Query().Get("delta"))
	if err != nil {
		http_utils.HttpCustomError(w, errors.New("invalid delta"))
		return
	}

	var dateFrom, dateTo models.Millis
	dateTo = models.Millis(time.Now())
	var measure models.Timespan

	switch measureQuery {
	case "minute":
		dateFrom = models.Millis(time.Now().Add(-1 * time.Minute * time.Duration(delta)))
		measure = models.Minute

	case "hour":
		dateFrom = models.Millis(time.Now().Add(-1 * time.Hour * time.Duration(delta)))
		measure = models.Hour

	case "day":
		dateFrom = models.Millis(time.Now().AddDate(0, 0, -1*delta))
		measure = models.Day

	case "week":
		dateFrom = models.Millis(time.Now().AddDate(0, 0, -7*delta))
		measure = models.Week

	case "month":
		dateFrom = models.Millis(time.Now().AddDate(0, -delta, 0))
		measure = models.Month

	case "year":
		dateFrom = models.Millis(time.Now().AddDate(-delta, 0, 0))
		measure = models.Year

	default:
		http_utils.HttpCustomError(w, errors.New("invalid measure"))
		return
	}

	client := lib.PolygonClient

	res := client.ListAggs(r.Context(), &models.ListAggsParams{
		Ticker:     symbol,
		Multiplier: multiplier,
		Timespan:   measure,
		From:       dateFrom,
		To:         dateTo,
	})

	results := []models.Agg{}

	for {
		results = append(results, res.Item())
		if !res.Next() {
			break
		}
	}

	render.JSON(w, r, results)
}

func ResearchSymbolNews(w http.ResponseWriter, r *http.Request) {
	symbol := chi.URLParam(r, "symbol")
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil || limit > 100 || limit < 0 {
		http_utils.HttpCustomError(w, errors.New("invalid limit"))
		return
	}

	client := lib.PolygonClient

	// today := time.Now()
	params := models.ListTickerNewsParams{
		TickerEQ: &symbol,
	}.WithLimit(limit)

	asyncIter := client.ListTickerNews(context.Background(), params, models.WithTrace(true))

	results := []models.TickerNews{}
	count := 0
	asyncIter.Next()
	for {
		results = append(results, asyncIter.Item())
		count++
		if count >= limit {
			break
		}
		if !asyncIter.Next() {
			break
		}
	}

	render.JSON(w, r, results)

}
