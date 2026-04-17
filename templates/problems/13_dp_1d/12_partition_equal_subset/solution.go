package partition_equal_subset

// solveCanPartition checks if the array can be partitioned into two equal-sum subsets. O(n * sum) time, O(sum) space.
func solveCanPartition(nums []int) bool {
	total := 0
	for _, n := range nums {
		total += n
	}
	if total%2 != 0 {
		return false
	}
	target := total / 2
	dp := make([]bool, target+1)
	dp[0] = true
	for _, num := range nums {
		for j := target; j >= num; j-- {
			dp[j] = dp[j] || dp[j-num]
		}
	}
	return dp[target]
}
