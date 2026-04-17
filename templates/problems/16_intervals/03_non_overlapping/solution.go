package non_overlapping

import "sort"

// solveEraseOverlapIntervals finds min removals using greedy interval scheduling. O(n log n) time, O(1) space.
func solveEraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count := 0
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}
	return count
}
