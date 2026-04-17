package single_number

// solveSingleNumber XORs all elements; duplicates cancel out. O(n) time, O(1) space.
func solveSingleNumber(nums []int) int {
	result := 0
	for _, n := range nums {
		result ^= n
	}
	return result
}
