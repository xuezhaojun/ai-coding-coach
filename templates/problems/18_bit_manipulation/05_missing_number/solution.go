package missing_number

// solveMissingNumber uses XOR to find the missing number. O(n) time, O(1) space.
func solveMissingNumber(nums []int) int {
	result := len(nums)
	for i, v := range nums {
		result ^= i ^ v
	}
	return result
}
