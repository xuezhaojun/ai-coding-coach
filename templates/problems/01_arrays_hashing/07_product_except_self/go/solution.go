package product_except_self

// solveProductExceptSelf computes product array without division using prefix/suffix passes.
// Time: O(n), Space: O(1) extra (output array not counted)
func solveProductExceptSelf(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	// Forward pass: result[i] = product of all elements before i
	result[0] = 1
	for i := 1; i < n; i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	// Backward pass: multiply by product of all elements after i
	suffix := 1
	for i := n - 2; i >= 0; i-- {
		suffix *= nums[i+1]
		result[i] *= suffix
	}

	return result
}
