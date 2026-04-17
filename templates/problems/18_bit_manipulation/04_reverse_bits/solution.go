package reverse_bits

// solveReverseBits reverses all 32 bits. O(1) time (32 iterations), O(1) space.
func solveReverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result = (result << 1) | (num & 1)
		num >>= 1
	}
	return result
}
