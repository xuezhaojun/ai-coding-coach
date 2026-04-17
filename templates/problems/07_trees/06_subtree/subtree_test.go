package subtree

import "testing"

func TestIsSubtree(t *testing.T) {
	tests := []struct {
		name    string
		root    []int
		subRoot []int
		want    bool
	}{
		{"is subtree", []int{3, 4, 5, 1, 2}, []int{4, 1, 2}, true},
		{"not subtree", []int{3, 4, 5, 1, 2, -101, -101, -101, -101, 0}, []int{4, 1, 2}, false},
		{"both single same", []int{1}, []int{1}, true},
		{"both single diff", []int{1}, []int{2}, false},
		{"sub is nil", []int{1, 2, 3}, []int{}, true},
		{"root equals sub", []int{1, 2, 3}, []int{1, 2, 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeHelper(tt.root)
			subRoot := buildTreeHelper(tt.subRoot)
			if got := isSubtree(root, subRoot); got != tt.want {
				t.Errorf("isSubtree() = %v, want %v", got, tt.want)
			}
		})
	}
}
