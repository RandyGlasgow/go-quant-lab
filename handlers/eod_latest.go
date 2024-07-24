package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/polygon-io/client-go/rest/models"
	"github.com/vec-search/lib"
	"github.com/vec-search/lib/http/http_utils"
	"github.com/vec-search/lib/market_stack/ms_utils"
)

func EndOfDayLatest(w http.ResponseWriter, r *http.Request) {
	c := lib.PolyIoClient.Client
	symbols, err := ms_utils.ExtractSymbolsFromQuery(r)
	if err != nil {
		http_utils.HttpCustomError(w, err)
		return
	}

	splitSymbols := strings.Split(symbols, ",")

	// Create a channel for results and errors
	resultsCh := make(chan models.GetDailyOpenCloseAggResponse)
	errorsCh := make(chan error, len(splitSymbols))
	unknownSymbolsCh := make(chan string, len(splitSymbols))

	// Launch goroutines for concurrent requests
	for _, symbol := range splitSymbols {
		if symbol == "" {
			continue
		}

		go func(sym string) {
			defer func() {
				if err := recover(); err != nil {
					errorsCh <- fmt.Errorf("error for symbol %s: %w", sym, errors.Unwrap(err.(error)))
				}
			}()

			// Set params within the goroutine
			params := &models.GetDailyOpenCloseAggParams{
				Ticker: sym,
				Date:   models.Date(time.Now()),
			}

			// Make request within the goroutine
			res, err := c.GetDailyOpenCloseAgg(context.Background(), params)
			if err != nil {
				if strings.Contains(err.Error(), "404") {
					unknownSymbolsCh <- sym
					return
				}

				errorsCh <- err
				return
			}

			resultsCh <- *res
		}(symbol)
	}

	// Collect results and errors concurrently
	resArr := []models.GetDailyOpenCloseAggResponse{}
	errArr := []error{}
	unknownSymbolsArr := []string{}

	for i := 0; i < len(splitSymbols); i++ {
		select {
		case res := <-resultsCh:
			resArr = append(resArr, res)
		case err := <-errorsCh:
			errArr = append(errArr, err)

		// Handle unknown symbols
		case sym := <-unknownSymbolsCh:
			unknownSymbolsArr = append(unknownSymbolsArr, sym)
		}
	}

	// close the channels
	close(resultsCh)
	close(errorsCh)
	close(unknownSymbolsCh)

	// Render the response
	render.JSON(w, r, map[string]interface{}{
		"results": resArr,
		"unknown": unknownSymbolsArr,
		"errors":  errArr,
	})
}

func EndOfDayByDate(w http.ResponseWriter, r *http.Request) {
	c := lib.PolyIoClient.Client
	symbols, err := ms_utils.ExtractSymbolsFromQuery(r)
	if err != nil {
		http_utils.HttpCustomError(w, err)
		return
	}

	date := chi.URLParam(r, "date")
	if date == "" {
		http_utils.HttpCustomError(w, errors.New("date is required"))
		return
	}

	splitSymbols := strings.Split(symbols, ",")

	// Create a channel for results and errors
	resultsCh := make(chan models.GetDailyOpenCloseAggResponse)
	errorsCh := make(chan error, len(splitSymbols))
	unknownSymbolsCh := make(chan string, len(splitSymbols))

	// Launch goroutines for concurrent requests
	for _, symbol := range splitSymbols {
		if symbol == "" {
			continue
		}

		go func(sym string) {
			defer func() {
				if err := recover(); err != nil {
					errorsCh <- fmt.Errorf("error for symbol %s: %w", sym, errors.Unwrap(err.(error)))
				}
			}()

			// Parse the date
			dateModes, err := time.Parse("2006-01-02", date)
			if err != nil {
				errorsCh <- err
				return
			}

			// Set params within the goroutine
			params := &models.GetDailyOpenCloseAggParams{
				Ticker: sym,
				Date:   models.Date(dateModes),
			}

			// Make request within the goroutine
			res, err := c.GetDailyOpenCloseAgg(context.Background(), params)
			if err != nil {
				if strings.Contains(err.Error(), "404") {
					unknownSymbolsCh <- sym
					return
				}

				errorsCh <- err
				return
			}

			resultsCh <- *res
		}(symbol)
	}

	// Collect results and errors concurrently
	resArr := []models.GetDailyOpenCloseAggResponse{}
	errArr := []error{}
	unknownSymbolsArr := []string{}

	for i := 0; i < len(splitSymbols); i++ {
		select {
		case res := <-resultsCh:
			resArr = append(resArr, res)
		case err := <-errorsCh:
			errArr = append(errArr, err)

		// Handle unknown symbols
		case sym := <-unknownSymbolsCh:
			unknownSymbolsArr = append(unknownSymbolsArr, sym)
		}
	}

	// close the channels
	close(resultsCh)
	close(errorsCh)
	close(unknownSymbolsCh)

	// Render the response
	render.JSON(w, r, map[string]interface{}{
		"results": resArr,
		"unknown": unknownSymbolsArr,
		"errors":  errArr,
	})
}
