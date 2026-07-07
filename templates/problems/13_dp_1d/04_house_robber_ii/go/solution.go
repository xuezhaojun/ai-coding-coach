package house_robber_ii

// solveRobII returns the maximum amount robbed from circular houses. O(n) time, O(1) space.
func solveRobII(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

func robRange(nums []int, lo, hi int) int {
	prev, curr := 0, 0
	for i := lo; i <= hi; i++ {
		prev, curr = curr, max(curr, prev+nums[i])
	}
	return curr
}
