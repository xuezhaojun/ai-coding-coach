package plus_one

// solvePlusOne adds one to a number represented as digits. O(n) time, O(1) space (or O(n) if carry propagates).
func solvePlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	// All digits were 9, need extra digit
	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}
