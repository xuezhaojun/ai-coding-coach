package insert_interval

// solveInsert merges a new interval into a sorted list. O(n) time, O(n) space.
func solveInsert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	i := 0
	// Add all intervals before the new interval
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}
	// Merge overlapping intervals with newInterval
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if intervals[i][1] > newInterval[1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	result = append(result, newInterval)
	// Add remaining intervals
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}
	return result
}
