package problems

func MaxProfit(prices []int) int {
	minPrice := prices[0]
	profit := 0
	sell := 0
	for i := 1; i < len(prices); i++ {
		sell = prices[i] - minPrice
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
		if sell > profit {
			profit = sell
		}
	}
	return profit
}
