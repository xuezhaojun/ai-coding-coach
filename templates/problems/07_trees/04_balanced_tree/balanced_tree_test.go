package balanced_tree

import "testing"

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want bool
	}{
		{"balanced", []int{3, 9, 20, -101, -101, 15, 7}, true},
		{"unbalanced", []int{1, 2, 2, 3, 3, -101, -101, 4, 4}, false},
		{"empty tree", []int{}, true},
		{"single node", []int{1}, true},
		{"two levels balanced", []int{1, 2, 3}, true},
		{"left heavy by one", []int{1, 2, -101, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeHelper(tt.vals)
			if got := isBalanced(root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
