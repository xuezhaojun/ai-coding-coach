package coin_change_ii

// solveChange returns the number of combinations to make the amount. O(amount * len(coins)) time, O(amount) space.
func solveChange(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			dp[j] += dp[j-coin]
		}
	}
	return dp[amount]
}
