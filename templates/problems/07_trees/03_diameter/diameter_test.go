package diameter

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{"example 1", []int{1, 2, 3, 4, 5}, 3},
		{"example 2", []int{1, 2}, 1},
		{"single node", []int{1}, 0},
		{"empty tree", []int{}, 0},
		{"left skewed", []int{1, 2, -101, 3, -101, 4}, 3},
		{"wide tree", []int{1, 2, 3, 4, 5, 6, 7}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeHelper(tt.vals)
			if got := diameterOfBinaryTree(root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
