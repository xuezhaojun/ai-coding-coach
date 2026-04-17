package best_time_stock

// solveMaxProfit finds max profit from one buy/sell using a single pass.
// Time: O(n), Space: O(1)
func solveMaxProfit(prices []int) int {
	minPrice := prices[0]
	maxProfit := 0
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else if p-minPrice > maxProfit {
			maxProfit = p - minPrice
		}
	}
	return maxProfit
}
