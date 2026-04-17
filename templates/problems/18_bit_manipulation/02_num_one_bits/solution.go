package num_one_bits

// solveHammingWeight counts set bits using n & (n-1) trick. O(k) time where k = number of 1 bits, O(1) space.
func solveHammingWeight(n uint32) int {
	count := 0
	for n != 0 {
		n &= n - 1
		count++
	}
	return count
}
