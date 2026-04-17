package meeting_rooms_ii

import "sort"

// solveMinMeetingRooms uses a sweep line approach. O(n log n) time, O(n) space.
func solveMinMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	starts := make([]int, len(intervals))
	ends := make([]int, len(intervals))
	for i, iv := range intervals {
		starts[i] = iv[0]
		ends[i] = iv[1]
	}
	sort.Ints(starts)
	sort.Ints(ends)

	rooms, maxRooms := 0, 0
	s, e := 0, 0
	for s < len(starts) {
		if starts[s] < ends[e] {
			rooms++
			s++
		} else {
			rooms--
			e++
		}
		if rooms > maxRooms {
			maxRooms = rooms
		}
	}
	return maxRooms
}
