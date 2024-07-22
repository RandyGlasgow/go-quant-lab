package ms_utils

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func ExtractSymbolsFromQuery(r *http.Request) (string, error) {
	if r.URL.Query().Has("symbols") == false {
		return "",
			errors.New("symbols query parameter is required")
	}

	// cant be greater than 100 symbols

	symbols := r.URL.Query().Get("symbols")

	// prepare symbols
	upperCase := strings.ToUpper(symbols)

	if len(strings.Split(upperCase, ",")) > 100 {
		return "",
			errors.New("symbols query parameter cannot be greater than 100 symbols")
	}

	fmt.Println("uppercase symbols:", upperCase)

	return upperCase, nil
}

func parseDate(date string) (string, error) {
	_, err := time.Parse(time.DateOnly, date)
	_, err2 := time.Parse("2006-01-02T15:04:05", date)
	if err != nil && err2 != nil {
		return "", errors.New("invalid date format")
	}

	return date, nil
}

func ExtractDateFromQuery(r *http.Request) (string, error) {
	q := r.URL.Query()

	if q.Has("date") {
		return parseDate(q.Get("date"))
	}

	return "", errors.New("date query parameter is required")
}

func ExtractOptionalToDateFromQuery(r *http.Request) (string, error) {
	q := r.URL.Query()
	if q.Has("date_to") {
		return parseDate(q.Get("date_to"))
	}
	return "", nil
}

func ExtractOptionalFromDateFromQuery(r *http.Request) (string, error) {
	q := r.URL.Query()
	if q.Has("date_from") {
		return parseDate(q.Get("date_from"))
	}
	return "", nil
}

func ExtractOptionalLimitFromQuery(r *http.Request) (string, error) {
	q := r.URL.Query()

	limit := q.Get("limit")

	if limit == "" {
		return "100", nil
	}
	// check that the limit is greater than 0 and less than 1000
	if limitInt, err := strconv.Atoi(limit); err != nil || limitInt < 0 || limitInt > 1000 {
		return "", errors.New("limit query parameter must be greater than 0 and less than 1000")
	}

	return limit, nil
}

func ExtractOffsetFromQuery(r *http.Request) (string, error) {
	q := r.URL.Query()

	offset := q.Get("offset")

	if offset == "" {
		return "0", nil
	}
	// check that the offset is greater than 0
	if offsetInt, err := strconv.Atoi(offset); err != nil || offsetInt < 0 {
		return "", errors.New("offset query parameter must be greater than 0")
	}

	return offset, nil
}

func ExtractSortFromQuery(r *http.Request) (string, error) {
	q := r.URL.Query()

	sort := q.Get("sort")

	if sort == "" {
		return "", errors.New("sort query parameter is required")
	}

	if sort != "asc" && sort != "desc" {
		return "", errors.New("sort query parameter must be either asc or desc")
	}

	return sort, nil
}

func ExtractOptionalExchangeFromQuery(r *http.Request) (string, error) {
	q := r.URL.Query()

	exchange := q.Get("exchange")

	if exchange == "" {
		return "", errors.New("exchange query parameter is required")
	}

	return exchange, nil
}
