package max_subarray

// solveMaxSubArray uses Kadane's algorithm. O(n) time, O(1) space.
func solveMaxSubArray(nums []int) int {
	maxSum := nums[0]
	curSum := nums[0]
	for i := 1; i < len(nums); i++ {
		if curSum < 0 {
			curSum = nums[i]
		} else {
			curSum += nums[i]
		}
		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}
