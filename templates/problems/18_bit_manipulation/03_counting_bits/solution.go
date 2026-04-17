package counting_bits

// solveCountBits uses DP with the relation dp[i] = dp[i>>1] + (i&1). O(n) time, O(n) space.
func solveCountBits(n int) []int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i>>1] + (i & 1)
	}
	return dp
}
