package unique_paths

// solveUniquePaths returns the number of unique paths in an m x n grid. O(m*n) time, O(n) space.
func solveUniquePaths(m int, n int) int {
	dp := make([]int, n)
	for j := range dp {
		dp[j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
