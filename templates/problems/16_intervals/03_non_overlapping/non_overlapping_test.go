package non_overlapping

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  int
	}{
		{"one removal", [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{"no removal", [][]int{{1, 2}, {2, 3}}, 0},
		{"all overlap", [][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{"single interval", [][]int{{1, 5}}, 0},
		{"nested intervals", [][]int{{1, 10}, {2, 3}, {4, 5}}, 1},
		{"chain overlap", [][]int{{1, 3}, {2, 4}, {3, 5}}, 1},
		{"negative start LeetCode edge", [][]int{{-50000, 1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := eraseOverlapIntervals(tt.intervals)
			if got != tt.expected {
				t.Errorf("eraseOverlapIntervals(%v) = %d, want %d", tt.intervals, got, tt.expected)
			}
		})
	}
}
