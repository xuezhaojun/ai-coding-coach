package reverse_integer

import "math"

// solveReverse reverses digits with overflow check. O(log x) time, O(1) space.
func solveReverse(x int) int {
	result := 0
	for x != 0 {
		digit := x % 10
		x /= 10
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			return 0
		}
		if result < math.MinInt32/10 || (result == math.MinInt32/10 && digit < -8) {
			return 0
		}
		result = result*10 + digit
	}
	return result
}
