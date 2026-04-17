package min_cost_stairs

// solveMinCostClimbingStairs returns the min cost to reach the top. O(n) time, O(1) space.
func solveMinCostClimbingStairs(cost []int) int {
	n := len(cost)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return cost[0]
	}
	prev, curr := cost[0], cost[1]
	for i := 2; i < n; i++ {
		prev, curr = curr, cost[i] + min(prev, curr)
	}
	return min(prev, curr)
}

func minVal(a, b int) int {
	if a < b {
		return a
	}
	return b
}
