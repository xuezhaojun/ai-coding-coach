package stock_cooldown

// solveMaxProfitCooldown returns the max profit with cooldown. O(n) time, O(1) space.
func solveMaxProfitCooldown(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	// held: max profit while holding a stock
	// sold: max profit on the day we just sold
	// rest: max profit while resting (cooldown or idle)
	held, sold, rest := -prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		prevSold := sold
		sold = held + prices[i]
		held = max(held, rest-prices[i])
		rest = max(rest, prevSold)
	}
	return max(sold, rest)
}
