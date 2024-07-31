package research

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/render"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/vec-search/lib"
	"github.com/vec-search/lib/http/http_utils"
)

func ResearchSymbolSnapshot(w http.ResponseWriter, r *http.Request) {
	symbols := r.URL.Query().Get("symbols")

	if symbols == "" {
		http_utils.HttpCustomError(w, errors.New("symbols is required"))
		return
	}

	// uppercase the symbols
	symbols = strings.ToUpper(symbols)

	client := lib.PolygonClient
	params := models.GetAllTickersSnapshotParams{
		Tickers:    &symbols,
		MarketType: models.Stocks,
		Locale:     models.US,
	}

	data, err := client.GetAllTickersSnapshot(r.Context(), &params)
	fmt.Print(err)
	if err != nil {
		http_utils.HttpStandardError(w, http.StatusBadRequest)
		return
	}

	// Get the ticker details
	render.JSON(w, r, data)
}
