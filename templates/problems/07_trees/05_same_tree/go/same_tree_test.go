package same_tree

import "testing"

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		name string
		p    []int
		q    []int
		want bool
	}{
		{"identical", []int{1, 2, 3}, []int{1, 2, 3}, true},
		{"different structure", []int{1, 2}, []int{1, -101, 2}, false},
		{"different values", []int{1, 2, 1}, []int{1, 1, 2}, false},
		{"both empty", []int{}, []int{}, true},
		{"one empty", []int{1}, []int{}, false},
		{"single same", []int{1}, []int{1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := buildTreeHelper(tt.p)
			q := buildTreeHelper(tt.q)
			if got := isSameTree(p, q); got != tt.want {
				t.Errorf("isSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
