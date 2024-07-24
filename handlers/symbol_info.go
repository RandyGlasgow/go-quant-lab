package handlers

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/vec-search/lib"
	"github.com/vec-search/lib/http/http_utils"
	"net/http"
	"time"
)

func SymbolInfo(w http.ResponseWriter, r *http.Request) {
	date := chi.URLParam(r, "date")

	if date == "" {
		http_utils.HttpCustomError(w, errors.New("date is required"))
		return
	}
	// Parse the date
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		http_utils.HttpCustomError(w, err)
		return
	}
	dateModes := models.Date(parsedDate)

	params := &models.GetTickerDetailsParams{
		Ticker: r.URL.Query().Get("symbol"),
	}
	params.WithDate(dateModes)

	res, err := lib.PolyIoClient.Client.GetTickerDetails(r.Context(), params)

	if err != nil {
		http_utils.HttpCustomError(w, err)
		return
	}

	render.JSON(w, r, res)
}
