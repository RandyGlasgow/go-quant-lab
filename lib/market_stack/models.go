package market_stack

type Pagination struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Count  int `json:"count"`
	Total  int `json:"total"`
}

type EndOfDayService struct {
	Pagination Pagination `json:"pagination"`
	Data       []struct {
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

/*
*
"name": "Microsoft Corporation",

	"symbol": "MSFT",
	"has_intraday": false,
	"has_eod": true,
	"country": null,
	"stock_exchange": {
	  "name": "NASDAQ Stock Exchange",
	  "acronym": "NASDAQ",
	  "mic": "XNAS",
	  "country": "USA",
	  "country_code": "US",
	  "city": "New York",
	  "website": "WWW.NASDAQ.COM"
	}
*/
type TickerInfo struct {
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	HasIntraday   bool   `json:"has_intraday"`
	HasEod        bool   `json:"has_eod"`
	Country       string `json:"country"`
	StockExchange struct {
		Name        string `json:"name"`
		Acronym     string `json:"acronym"`
		Mic         string `json:"mic"`
		Country     string `json:"country"`
		CountryCode string `json:"country_code"`
		City        string `json:"city"`
	}
}
