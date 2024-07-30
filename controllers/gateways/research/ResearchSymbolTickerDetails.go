package research

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/vec-search/lib"
	"github.com/vec-search/lib/http/http_utils"
)

func ResearchSymbolTickerDetails(w http.ResponseWriter, r *http.Request) {
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
