package stock

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"time"
)

const (
	AlphaVantageAPIKey = "FTLH2B4CW8XOBF9O"
	DateLayout         = "2006-01-02"
	StockDataURL       = "https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&outputsize=full&apikey=%s"
)

type AlphaVantageData struct {
	MetaDataInfo     MetaData         `json:"Meta Data"`
	StockPriceSeries map[string]Price `json:"Time Series (Daily)"`
}

type MetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	OutputSize    string `json:"4. Output Size"`
	TimeZone      string `json:"5. Time Zone"`
}

type Price struct {
	OpenPrice  string `json:"1. open"`
	HighPrice  string `json:"2. high"`
	LowPrice   string `json:"3. low"`
	ClosePrice string `json:"4. close"`
	Volume     string `json:"5. volume"`
}

type StockPrice struct {
	StockDate string
	HighPrice float64
}

type timeSlice []time.Time

func (s timeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s timeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s timeSlice) Len() int           { return len(s) }

var StockPrices []StockPrice

func GetData(symbol string) ([]StockPrice, error) {

	url := fmt.Sprintf(StockDataURL, symbol, AlphaVantageAPIKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	dataObject := &AlphaVantageData{}
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(dataObject)
	if err != nil {
		return nil, err
	}

	dataPoints := len(dataObject.StockPriceSeries)
	log.Println("total data points:")
	log.Println(dataPoints)

	if dataPoints < 2 {
		return nil, errors.New("either the input symbol is invalid or insufficient history")
	}

	keys := make(timeSlice, dataPoints)
	index := 0
	for key := range dataObject.StockPriceSeries {
		t, err := time.Parse(DateLayout, key)
		if err != nil {
			return nil, err
		}

		keys[index] = t
		index++
	}

	sort.Sort(keys)

	stockPrices := make([]StockPrice, dataPoints)
	indx := 0
	for _, key := range keys {
		f, err := strconv.ParseFloat(dataObject.StockPriceSeries[key.Format(DateLayout)].HighPrice, 64)
		if err != nil {
			return nil, errors.New("failed to process stock data")
		}
		stockPrices[indx] = StockPrice{StockDate: key.Format(DateLayout), HighPrice: f}
		indx++
	}

	return stockPrices, nil
}
