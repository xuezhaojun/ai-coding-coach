package distinct_subsequences

// solveNumDistinct returns the number of distinct subsequences of s that equal t. O(m*n) time, O(n) space.
func solveNumDistinct(s string, t string) int {
	m, n := len(s), len(t)
	if m < n {
		return 0
	}
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= m; i++ {
		// iterate backwards to avoid overwriting values we still need
		for j := min(i, n); j >= 1; j-- {
			if s[i-1] == t[j-1] {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n]
}
