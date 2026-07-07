package meeting_rooms_ii

import "testing"

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{"two rooms", [][]int{{0, 30}, {5, 10}, {15, 20}}, 2},
		{"one room", [][]int{{7, 10}, {2, 4}}, 1},
		{"all overlap", [][]int{{1, 5}, {2, 6}, {3, 7}}, 3},
		{"single meeting", [][]int{{1, 10}}, 1},
		{"sequential", [][]int{{0, 5}, {5, 10}, {10, 15}}, 1},
		{"nested", [][]int{{1, 10}, {2, 5}, {6, 9}}, 2},
		{"empty", [][]int{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minMeetingRooms(tt.intervals)
			if got != tt.expected {
				t.Errorf("minMeetingRooms(%v) = %d, want %d", tt.intervals, got, tt.expected)
			}
		})
	}
}
