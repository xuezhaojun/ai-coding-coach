package max_path_sum

import "testing"

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{"example 1", []int{1, 2, 3}, 6},
		{"example 2", []int{-10, 9, 20, -101, -101, 15, 7}, 42},
		{"single node", []int{1}, 1},
		{"single negative", []int{-3}, -3},
		{"all negative", []int{-1, -2, -3}, -1},
		{"mixed", []int{5, 4, 8, 11, -101, 13, 4, 7, 2, -101, -101, -101, 1}, 48},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeHelper(tt.vals)
			if got := maxPathSum(root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
