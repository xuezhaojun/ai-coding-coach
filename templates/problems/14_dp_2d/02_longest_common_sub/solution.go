package longest_common_sub

// solveLongestCommonSubsequence returns the length of the LCS. O(m*n) time, O(min(m,n)) space.
func solveLongestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) < len(text2) {
		text1, text2 = text2, text1
	}
	m, n := len(text1), len(text2)
	dp := make([]int, n+1)
	for i := 1; i <= m; i++ {
		prev := 0
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if text1[i-1] == text2[j-1] {
				dp[j] = prev + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			prev = tmp
		}
	}
	return dp[n]
}
