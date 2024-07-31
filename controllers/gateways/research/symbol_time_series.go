package research

import (
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
