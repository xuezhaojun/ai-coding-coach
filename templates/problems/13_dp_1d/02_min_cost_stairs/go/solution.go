package min_cost_stairs

// solveMinCostClimbingStairs returns the min cost to reach the top. O(n) time, O(n) space.
// dp[i] = min cost to reach step i (not yet paying cost[i])
func solveMinCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0] = 0 // free to start at step 0
	dp[1] = 0 // free to start at step 1
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}
