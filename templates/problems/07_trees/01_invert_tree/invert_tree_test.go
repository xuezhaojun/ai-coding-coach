package invert_tree

import (
	"reflect"
	"testing"
)

func TestInvertTree(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"example 1", []int{4, 2, 7, 1, 3, 6, 9}, []int{4, 7, 2, 9, 6, 3, 1}},
		{"example 2", []int{2, 1, 3}, []int{2, 3, 1}},
		{"empty tree", []int{}, []int{}},
		{"single node", []int{1}, []int{1}},
		{"left only", []int{1, 2}, []int{1, -101, 2}},
		{"right only", []int{1, -101, 2}, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeHelper(tt.vals)
			got := treeToSlice(invertTree(root))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
