package sum_two_integers

// solveGetSum adds two integers using bit manipulation. O(1) time (bounded iterations), O(1) space.
func solveGetSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}
