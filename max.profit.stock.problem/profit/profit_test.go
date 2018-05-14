package profit

import (
	"math/rand"
	"problem.solving/max.profit.stock.problem/stock"
	"testing"
	"time"
)

func populateRandomData(min float64, max float64, dataSize int) []stock.StockPrice {
	date := time.Now()
	stock_prices := make([]stock.StockPrice, dataSize)
	for i := 0; i < dataSize; i++ {
		price := min + rand.Float64()*(max-min)
		stock_prices[i] = stock.StockPrice{date.Format(stock.DateLayout), price}
		date = date.Add(time.Hour * 24)
	}
	return stock_prices
}

func populateKnownData() []stock.StockPrice {
	stock_prices := []stock.StockPrice{
		stock.StockPrice{time.Now().Format(stock.DateLayout), 100.0},
		stock.StockPrice{time.Now().Add(time.Hour * 24).Format(stock.DateLayout), 50.0},
		stock.StockPrice{time.Now().Add(time.Hour * 2 * 24).Format(stock.DateLayout), 150.0},
		stock.StockPrice{time.Now().Add(time.Hour * 3 * 24).Format(stock.DateLayout), 200.0},
		stock.StockPrice{time.Now().Add(time.Hour * 4 * 24).Format(stock.DateLayout), 10.0},
		stock.StockPrice{time.Now().Add(time.Hour * 5 * 24).Format(stock.DateLayout), 110.0},
		stock.StockPrice{time.Now().Add(time.Hour * 6 * 24).Format(stock.DateLayout), 500.0},
		stock.StockPrice{time.Now().Add(time.Hour * 7 * 24).Format(stock.DateLayout), 60.0},
		stock.StockPrice{time.Now().Add(time.Hour * 8 * 24).Format(stock.DateLayout), 250.0},
		stock.StockPrice{time.Now().Add(time.Hour * 9 * 24).Format(stock.DateLayout), 70.0},
	}
	return stock_prices
}

func populateSortedWorstData() []stock.StockPrice {
	stock_prices := []stock.StockPrice{
		stock.StockPrice{time.Now().Format(stock.DateLayout), 110.0},
		stock.StockPrice{time.Now().Add(time.Hour * 24).Format(stock.DateLayout), 100.0},
		stock.StockPrice{time.Now().Add(time.Hour * 2 * 24).Format(stock.DateLayout), 90.0},
		stock.StockPrice{time.Now().Add(time.Hour * 3 * 24).Format(stock.DateLayout), 80.0},
		stock.StockPrice{time.Now().Add(time.Hour * 4 * 24).Format(stock.DateLayout), 70.0},
		stock.StockPrice{time.Now().Add(time.Hour * 5 * 24).Format(stock.DateLayout), 60.0},
		stock.StockPrice{time.Now().Add(time.Hour * 6 * 24).Format(stock.DateLayout), 50.0},
		stock.StockPrice{time.Now().Add(time.Hour * 7 * 24).Format(stock.DateLayout), 40.0},
		stock.StockPrice{time.Now().Add(time.Hour * 8 * 24).Format(stock.DateLayout), 30.0},
		stock.StockPrice{time.Now().Add(time.Hour * 9 * 24).Format(stock.DateLayout), 20.0},
	}
	return stock_prices
}

func populateSortedBestData() []stock.StockPrice {
	stock_prices := []stock.StockPrice{
		stock.StockPrice{time.Now().Format(stock.DateLayout), 20.0},
		stock.StockPrice{time.Now().Add(time.Hour * 24).Format(stock.DateLayout), 30.0},
		stock.StockPrice{time.Now().Add(time.Hour * 2 * 24).Format(stock.DateLayout), 40.0},
		stock.StockPrice{time.Now().Add(time.Hour * 3 * 24).Format(stock.DateLayout), 50.0},
		stock.StockPrice{time.Now().Add(time.Hour * 4 * 24).Format(stock.DateLayout), 60.0},
		stock.StockPrice{time.Now().Add(time.Hour * 5 * 24).Format(stock.DateLayout), 70.0},
		stock.StockPrice{time.Now().Add(time.Hour * 6 * 24).Format(stock.DateLayout), 80.0},
		stock.StockPrice{time.Now().Add(time.Hour * 7 * 24).Format(stock.DateLayout), 90.0},
		stock.StockPrice{time.Now().Add(time.Hour * 8 * 24).Format(stock.DateLayout), 100.0},
		stock.StockPrice{time.Now().Add(time.Hour * 9 * 24).Format(stock.DateLayout), 110.0},
	}
	return stock_prices
}

func TestStockWithRandomizedInput(t *testing.T) {
	prices := populateRandomData(10.0, 100.0, 10)
	for i := 0; i < 10; i++ {
		t.Logf("date: %s >> price: %f", prices[i].StockDate, prices[i].HighPrice)
	}
	buy, sell, profit, err := FindMaxProfit(prices)
	if err != nil {
		t.Fail()
	}
	t.Logf("buy: %s, sell: %s, profit: %f", prices[buy].StockDate, prices[sell].StockDate, profit)
}

func TestWithIncompleteData(t *testing.T) {
	prices := populateRandomData(10.0, 100.0, 1)
	_, _, _, err := FindMaxProfit(prices)
	if err == nil {
		t.Fail()
	}
}

func TestDeterministicOutput(t *testing.T) {
	prices := populateKnownData()
	for i := 0; i < 10; i++ {
		t.Logf("date: %s >> price: %f", prices[i].StockDate, prices[i].HighPrice)
	}
	buy, sell, profit, err := FindMaxProfit(prices)
	if err != nil {
		t.Fail()
	}
	t.Logf("buy: %s, sell: %s, profit: %f", prices[buy].StockDate, prices[sell].StockDate, profit)
	if profit != 490.000000 {
		t.Fail()
	}
}

func TestWorstCaseOutput(t *testing.T) {
	prices := populateSortedWorstData()
	for i := 0; i < 10; i++ {
		t.Logf("date: %s >> price: %f", prices[i].StockDate, prices[i].HighPrice)
	}
	_, _, _, err := FindMaxProfit(prices)
	if err == nil {
		t.Fail()
	}
}

func TestBestCaseOutput(t *testing.T) {
	prices := populateSortedBestData()
	for i := 0; i < 10; i++ {
		t.Logf("date: %s >> price: %f", prices[i].StockDate, prices[i].HighPrice)
	}
	buy, sell, profit, err := FindMaxProfit(prices)
	if err != nil {
		t.Fail()
	}
	t.Logf("buy: %s, sell: %s, profit: %f", prices[buy].StockDate, prices[sell].StockDate, profit)
	if profit != 90.000000 {
		t.Fail()
	}
}
