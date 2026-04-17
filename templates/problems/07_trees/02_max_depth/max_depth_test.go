package max_depth

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{"example 1", []int{3, 9, 20, -101, -101, 15, 7}, 3},
		{"example 2", []int{1, -101, 2}, 2},
		{"empty tree", []int{}, 0},
		{"single node", []int{1}, 1},
		{"left skewed", []int{1, 2, -101, 3}, 3},
		{"balanced depth 4", []int{1, 2, 3, 4, 5, 6, 7, 8}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeHelper(tt.vals)
			if got := maxDepth(root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
