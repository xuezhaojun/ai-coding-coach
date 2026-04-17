package word_break

// solveWordBreak checks if s can be segmented into dictionary words. O(n^2) time, O(n) space.
func solveWordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		wordSet[w] = true
	}
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
