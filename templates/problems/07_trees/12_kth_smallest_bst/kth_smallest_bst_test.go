package kth_smallest_bst

import "testing"

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		k    int
		want int
	}{
		{"example 1", []int{3, 1, 4, -101, 2}, 1, 1},
		{"example 2", []int{5, 3, 6, 2, 4, -101, -101, 1}, 3, 3},
		{"single node k=1", []int{1}, 1, 1},
		{"k equals size", []int{2, 1, 3}, 3, 3},
		{"middle element", []int{3, 1, 4, -101, 2}, 2, 2},
		{"large bst k=4", []int{5, 3, 7, 1, 4, 6, 8}, 4, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeHelper(tt.vals)
			if got := kthSmallest(root, tt.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
