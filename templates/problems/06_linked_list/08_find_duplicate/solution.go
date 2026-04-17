package find_duplicate

// solveFindDuplicate finds the duplicate using Floyd's cycle detection. O(n) time, O(1) space.
func solveFindDuplicate(nums []int) int {
	// Phase 1: Find intersection point
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	// Phase 2: Find entrance to the cycle
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
