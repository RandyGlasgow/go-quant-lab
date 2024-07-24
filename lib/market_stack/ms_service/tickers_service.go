package ms_service

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/vec-search/lib/market_stack/ms_utils"
	"net/http"
	"strings"
)

func MS_GetSymbolInfo(r *http.Request) (*http.Response, error) {
	url := ms_utils.MarketStackRequestBuilder()
	symbol := chi.URLParam(r, "symbol") // update the url path to include the eod/latest endpoint
	fmt.Println("symbol:", symbol)
	url.Path += "tickers/" + strings.ToUpper(symbol)

	return http.Get(url.String())
}
