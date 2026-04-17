package climbing_stairs

// solveClimbStairs returns the number of ways to climb n stairs. O(n) time, O(1) space.
func solveClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev, curr := 1, 2
	for i := 3; i <= n; i++ {
		prev, curr = curr, prev+curr
	}
	return curr
}
