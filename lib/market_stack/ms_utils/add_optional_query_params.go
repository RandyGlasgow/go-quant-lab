package ms_utils

import (
	"net/http"
	"net/url"
)

func AddOptionalQueryParams(q *url.Values, r *http.Request) (url.Values, error) {
	if r.URL.Query().Has("limit") {
		limit, _ := ExtractOptionalLimitFromQuery(r)
		q.Add("limit", limit)
	}

	if r.URL.Query().Has("offset") {
		offset, _ := ExtractOffsetFromQuery(r)
		q.Add("offset", offset)
	}

	if r.URL.Query().Has("sort") {
		sort, _ := ExtractSortFromQuery(r)
		q.Add("sort", sort)
	}

	if r.URL.Query().Has("exchange") {
		exchange, _ := ExtractOptionalExchangeFromQuery(r)
		q.Add("exchange", exchange)
	}

	if r.URL.Query().Has("date_from") {
		date_from, _ := ExtractOptionalFromDateFromQuery(r)
		q.Add("date_from", date_from)
	}

	if r.URL.Query().Has("date_to") {
		date_to, _ := ExtractOptionalToDateFromQuery(r)
		q.Add("date_to", date_to)
	}

	return *q, nil
}
