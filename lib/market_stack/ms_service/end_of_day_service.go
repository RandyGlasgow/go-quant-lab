package ms_service

import (
	"log"
	"net/http"

	"github.com/vec-search/lib/market_stack/ms_utils"
)

func MS_GetSymbolsEndOfDayLatest(r *http.Request) (*http.Response, error) {

	url := ms_utils.MarketStackRequestBuilder()
	symbols, err := ms_utils.ExtractSymbolsFromQuery(r)

	if err != nil {
		return nil, err
	}
	// update the url path to include the eod/latest endpoint
	url.Path += "eod/latest"

	// add the symbols query parameter
	q := url.Query()
	q.Add("symbols", symbols)
	ms_utils.AddOptionalQueryParams(&q, r)
	url.RawQuery = q.Encode()

	log.Println(url.String())
	return http.Get(url.String())
}

func MS_GetSymbolsEndOfDayByDate(r *http.Request) (*http.Response, error) {
	url := ms_utils.MarketStackRequestBuilder()

	symbols, err := ms_utils.ExtractSymbolsFromQuery(r)
	if err != nil {
		return nil, err
	}

	date, err := ms_utils.ExtractDateFromQuery(r)

	if err != nil {
		return nil, err
	}

	// update the url path to include the eod/latest endpoint
	url.Path += "eod/" + date

	// add the symbols query parameter
	q := url.Query()
	q.Add("symbols", symbols)
	url.RawQuery = q.Encode()

	return http.Get(url.String())
}
