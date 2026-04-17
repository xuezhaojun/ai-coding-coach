package house_robber

// solveRob returns the maximum amount that can be robbed from non-adjacent houses. O(n) time, O(1) space.
func solveRob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	prev, curr := 0, 0
	for _, num := range nums {
		prev, curr = curr, max(curr, prev+num)
	}
	return curr
}
