package profit

import (
	"errors"
	"problem.solving/max.profit.stock.problem/stock"
)

func FindMaxProfit(stock_prices []stock.StockPrice) (int, int, float64, error) {

	if len(stock_prices) < 2 {
		return -1, -1, -1, errors.New("Not enough data points")
	}
	cur_index := len(stock_prices) - 1
	current_sell_date := cur_index
	current_buy_date := 0
	current_profit := stock_prices[current_sell_date].HighPrice - stock_prices[current_buy_date].HighPrice
	for cur_index >= 0 {

		//we found a better price to sell than the current sell price
		if stock_prices[current_sell_date].HighPrice < stock_prices[cur_index].HighPrice {
			current_sell_date = cur_index
		}

		//we found a price to buy-at which is smaller than the current sell price and hece will make some profit
		//but compare the current profit with the recorded profit and only then change the buy date
		if stock_prices[current_sell_date].HighPrice >= stock_prices[cur_index].HighPrice && current_sell_date > cur_index {
			if current_profit < (stock_prices[current_sell_date].HighPrice - stock_prices[cur_index].HighPrice) {
				current_profit = stock_prices[current_sell_date].HighPrice - stock_prices[cur_index].HighPrice
				current_buy_date = cur_index
			}
		}

		cur_index--
	}

	if current_buy_date == current_sell_date {
		return -1, -1, 0, errors.New("profit cannot be made")
	}

	return current_buy_date, current_sell_date, current_profit, nil
}
