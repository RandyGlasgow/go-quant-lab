package research

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/vec-search/lib"
	"github.com/vec-search/lib/http/http_utils"
)

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
