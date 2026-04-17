package meeting_rooms

import "testing"

func TestCanAttendMeetings(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  bool
	}{
		{"overlapping", [][]int{{0, 30}, {5, 10}, {15, 20}}, false},
		{"no overlap", [][]int{{7, 10}, {2, 4}}, true},
		{"empty", [][]int{}, true},
		{"single meeting", [][]int{{1, 5}}, true},
		{"touching boundaries", [][]int{{1, 5}, {5, 10}}, true},
		{"same time", [][]int{{1, 5}, {1, 5}}, false},
		{"sequential", [][]int{{0, 1}, {1, 2}, {2, 3}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canAttendMeetings(tt.intervals)
			if got != tt.expected {
				t.Errorf("canAttendMeetings(%v) = %v, want %v", tt.intervals, got, tt.expected)
			}
		})
	}
}
