package gas_station

// solveCanCompleteCircuit finds the starting gas station for a circular route. O(n) time, O(1) space.
func solveCanCompleteCircuit(gas []int, cost []int) int {
	totalSurplus := 0
	currentSurplus := 0
	start := 0
	for i := 0; i < len(gas); i++ {
		totalSurplus += gas[i] - cost[i]
		currentSurplus += gas[i] - cost[i]
		if currentSurplus < 0 {
			start = i + 1
			currentSurplus = 0
		}
	}
	if totalSurplus < 0 {
		return -1
	}
	return start
}
