package search_2d_matrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{"found in middle row", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{"not found", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
		{"found first element", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}}, 1, true},
		{"found last element", [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}}, 20, true},
		{"single element found", [][]int{{1}}, 1, true},
		{"single element not found", [][]int{{1}}, 2, false},
		{"single row", [][]int{{1, 3, 5}}, 3, true},
		{"single column", [][]int{{1}, {3}, {5}}, 5, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.matrix, tt.target); got != tt.want {
				t.Errorf("searchMatrix(%v, %d) = %v, want %v", tt.matrix, tt.target, got, tt.want)
			}
		})
	}
}
