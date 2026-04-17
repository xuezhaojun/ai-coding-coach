package search_rotated

import "testing"

func TestSearchRotated(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"found in left half", []int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{"found in right half", []int{4, 5, 6, 7, 0, 1, 2}, 5, 1},
		{"not found", []int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{"single element found", []int{1}, 1, 0},
		{"single element not found", []int{1}, 0, -1},
		{"not rotated found", []int{1, 2, 3, 4, 5}, 3, 2},
		{"two elements", []int{3, 1}, 1, 1},
		{"target at pivot", []int{6, 7, 1, 2, 3, 4, 5}, 1, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRotated(tt.nums, tt.target); got != tt.want {
				t.Errorf("searchRotated(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
