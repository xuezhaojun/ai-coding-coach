package target_sum

// solveFindTargetSumWays counts ways to assign +/- to reach target. O(n * sum) time, O(sum) space.
func solveFindTargetSumWays(nums []int, target int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	// We need (total + target) to be even and non-negative
	if (total+target)%2 != 0 || total+target < 0 || total < abs(target) {
		return 0
	}
	subsetSum := (total + target) / 2
	dp := make([]int, subsetSum+1)
	dp[0] = 1
	for _, num := range nums {
		for j := subsetSum; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}
	return dp[subsetSum]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
