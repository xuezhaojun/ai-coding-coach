package max_product_subarray

// solveMaxProduct returns the maximum product of a contiguous subarray. O(n) time, O(1) space.
func solveMaxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := nums[0]
	curMax, curMin := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			curMax, curMin = curMin, curMax
		}
		curMax = max(nums[i], curMax*nums[i])
		curMin = min(nums[i], curMin*nums[i])
		result = max(result, curMax)
	}
	return result
}
