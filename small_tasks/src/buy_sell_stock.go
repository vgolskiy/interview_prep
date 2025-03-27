package src

func MaxProfit(prices []int) int {
	var max_profit int
	if len(prices) > 1 {
		min_price_idx, max_price_idx := 0, 1
		for max_price_idx < len(prices) {
			if prices[max_price_idx] > prices[min_price_idx] {
				current_profit := prices[max_price_idx] - prices[min_price_idx]
				max_profit = max(current_profit, max_profit)
			} else {
				min_price_idx = max_price_idx
			}
			max_price_idx++
		}
	}
	return max_profit
}
