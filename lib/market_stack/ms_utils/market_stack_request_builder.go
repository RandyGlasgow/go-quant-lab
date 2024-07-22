package ms_utils

import (
	"errors"
	"log"
	"net/url"
	"os"
)

func MarketStackRequestBuilder() *url.URL {
	baseUrl := "https://api.marketstack.com/v1/"
	marketStackKey := os.Getenv("MARKET_STACK_API_KEY")

	if marketStackKey == "" {
		log.Fatal(errors.New("MarketStack API key not set"))
	}

	parsedUrl, _ := url.Parse(baseUrl)

	query := parsedUrl.Query()
	query.Set("access_key", marketStackKey)

	parsedUrl.RawQuery = query.Encode()

	log.Println(parsedUrl)

	return parsedUrl
}
