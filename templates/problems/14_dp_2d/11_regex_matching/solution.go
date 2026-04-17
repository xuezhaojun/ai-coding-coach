package regex_matching

// solveIsMatch checks if string s matches pattern p with '.' and '*'. O(m*n) time, O(n) space.
func solveIsMatch(s string, p string) bool {
	m, n := len(s), len(p)
	// dp[j] = whether s[0:i] matches p[0:j]
	dp := make([]bool, n+1)
	dp[0] = true
	// Initialize: patterns like a*, a*b*, a*b*c* can match empty string
	for j := 2; j <= n; j += 2 {
		if p[j-1] == '*' {
			dp[j] = dp[j-2]
		}
	}
	for i := 1; i <= m; i++ {
		prev := dp[0]
		dp[0] = false
		for j := 1; j <= n; j++ {
			tmp := dp[j]
			if p[j-1] == '*' {
				// zero occurrences of p[j-2]
				dp[j] = dp[j-2]
				// one or more occurrences
				if p[j-2] == '.' || p[j-2] == s[i-1] {
					dp[j] = dp[j] || tmp // tmp is dp[j] from previous row (i-1)
				}
			} else if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[j] = prev
			} else {
				dp[j] = false
			}
			prev = tmp
		}
	}
	return dp[n]
}
