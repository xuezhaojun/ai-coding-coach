package meeting_rooms

import "sort"

// solveCanAttendMeetings checks for overlaps after sorting. O(n log n) time, O(1) space.
func solveCanAttendMeetings(intervals [][]int) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < intervals[i-1][1] {
			return false
		}
	}
	return true
}
