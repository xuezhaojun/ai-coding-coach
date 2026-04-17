package validate_bst

import "testing"

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want bool
	}{
		{"valid bst", []int{2, 1, 3}, true},
		{"invalid bst", []int{5, 1, 4, -101, -101, 3, 6}, false},
		{"single node", []int{1}, true},
		{"empty tree", []int{}, true},
		{"left equal", []int{2, 2, 3}, false},
		{"valid larger", []int{5, 3, 7, 1, 4, 6, 8}, true},
		{"invalid deep right", []int{5, 4, 6, -101, -101, 3, 7}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeHelper(tt.vals)
			if got := isValidBST(root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
