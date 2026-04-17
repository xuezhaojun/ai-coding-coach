package largest_rect_histogram

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{"example 1", []int{2, 1, 5, 6, 2, 3}, 10},
		{"single bar", []int{5}, 5},
		{"increasing", []int{1, 2, 3, 4, 5}, 9},
		{"decreasing", []int{5, 4, 3, 2, 1}, 9},
		{"all same", []int{3, 3, 3, 3}, 12},
		{"two bars", []int{2, 4}, 4},
		{"valley", []int{6, 2, 5, 4, 5, 1, 6}, 12},
		{"empty", []int{}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.heights); got != tt.want {
				t.Errorf("largestRectangleArea(%v) = %d, want %d", tt.heights, got, tt.want)
			}
		})
	}
}
