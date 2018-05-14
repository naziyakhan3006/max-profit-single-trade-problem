package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"errors"
	"problem.solving/max.profit.stock.problem/profit"
	"problem.solving/max.profit.stock.problem/stock"
)

func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	body := []byte(fmt.Sprintf(`{"error":%q}`, err))
	w.Write(body)
}

func ProcessSymbol(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	symbol := params.ByName("symbol")
	days, err := strconv.Atoi(params.ByName("days"))
	if err != nil {
		handleError(err, w)
		log.Fatal("incorrect value specified for parameter: days")
		return
	}

	if days < 2 {
		handleError(errors.New("number of days must be two or more"), w)
		return
	}

	//Initialize the data
	stock_prices, err := stock.GetData(symbol)
	if err != nil {
		handleError(err, w)
		return
	}

	if len(stock_prices) > days {
		for _, k := range stock_prices[len(stock_prices)-days+1 : len(stock_prices)-1] {
			p := fmt.Sprintf(">> %s : %f", k.StockDate, k.HighPrice)
			log.Println(p)
		}
		stock_prices = stock_prices[len(stock_prices)-days+1 : len(stock_prices)-1]
	} else {
		log.Println("using the complete data set")
		log.Print("total data points considered:")
		log.Println(len(stock_prices))
	}

	buy, sell, profit, err := profit.FindMaxProfit(stock_prices)
	if err != nil {
		handleError(err, w)
		return
	}
	log.Println(">> Buy on : ", stock_prices[buy].StockDate)
	log.Println(">> Sell on : ", stock_prices[sell].StockDate)
	log.Println(">> Max Profit : ", profit)

	type Trade struct {
		BuyOn  string  `json:"buy on"`
		SellOn string  `json:"sell on"`
		Profit float64 `json:"profit"`
	}

	body, err := json.Marshal(Trade{BuyOn: stock_prices[buy].StockDate,
		SellOn: stock_prices[sell].StockDate,
		Profit: profit})
	if err != nil {
		handleError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	if body != nil {
		w.Write(body)
		w.Write([]byte("\n"))
	}
	return
}
