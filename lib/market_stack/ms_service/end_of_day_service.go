package ms_service

import (
	"log"
	"net/http"

	"github.com/vec-search/lib/market_stack/ms_utils"
)

type EndOfDayService struct {
	Pagination struct {
		Limit  int `json:"limit"`
		Offset int `json:"offset"`
		Count  int `json:"count"`
		Total  int `json:"total"`
	} `json:"pagination"`
	Data []struct {
		Open         float64 `json:"open"`
		High         float64 `json:"high"`
		Low          float64 `json:"low"`
		Close        float64 `json:"close"`
		Volume       float64 `json:"volume"`
		Adj_High     float64 `json:"adj_high"`
		Adj_Low      float64 `json:"adj_low"`
		Adj_Close    float64 `json:"adj_close"`
		Adj_Open     float64 `json:"adj_open"`
		Adj_Volume   float64 `json:"adj_volume"`
		Split_Factor float64 `json:"split_factor"`
		Dividend     float64 `json:"dividend"`
		Symbol       string  `json:"symbol"`
		Exchange     string  `json:"exchange"`
		Date         string  `json:"date"`
	} `json:"data"`
}

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
