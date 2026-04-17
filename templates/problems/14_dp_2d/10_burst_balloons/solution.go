package burst_balloons

// solveMaxCoins returns the maximum coins from bursting balloons. O(n^3) time, O(n^2) space.
func solveMaxCoins(nums []int) int {
	n := len(nums)
	// Add boundary balloons with value 1
	vals := make([]int, n+2)
	vals[0] = 1
	vals[n+1] = 1
	copy(vals[1:], nums)

	dp := make([][]int, n+2)
	for i := range dp {
		dp[i] = make([]int, n+2)
	}

	// length of the interval
	for length := 1; length <= n; length++ {
		for left := 1; left <= n-length+1; left++ {
			right := left + length - 1
			for k := left; k <= right; k++ {
				// k is the last balloon to burst in [left, right]
				coins := vals[left-1] * vals[k] * vals[right+1]
				coins += dp[left][k-1] + dp[k+1][right]
				if coins > dp[left][right] {
					dp[left][right] = coins
				}
			}
		}
	}
	return dp[1][n]
}
