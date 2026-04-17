package edit_distance

// solveMinDistance returns the minimum edit distance between two words. O(m*n) time, O(n) space.
func solveMinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([]int, n+1)
	for j := 0; j <= n; j++ {
		dp[j] = j
	}
	for i := 1; i <= m; i++ {
		prev := dp[0]
		dp[0] = i
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = prev
			} else {
				dp[j] = 1 + min(prev, min(dp[j], dp[j-1]))
			}
			prev = tmp
		}
	}
	return dp[n]
}
