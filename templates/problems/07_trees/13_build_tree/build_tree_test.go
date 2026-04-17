package build_tree

import (
	"reflect"
	"testing"
)

func TestBuildTreeFromPreIn(t *testing.T) {
	tests := []struct {
		name     string
		preorder []int
		inorder  []int
		want     []int
	}{
		{"example 1", []int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}, []int{3, 9, 20, -101, -101, 15, 7}},
		{"single node", []int{1}, []int{1}, []int{1}},
		{"left only", []int{1, 2}, []int{2, 1}, []int{1, 2}},
		{"right only", []int{1, 2}, []int{1, 2}, []int{1, -101, 2}},
		{"empty", []int{}, []int{}, []int{}},
		{"three nodes", []int{1, 2, 3}, []int{2, 1, 3}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := treeToSlice(buildTreeFromPreIn(tt.preorder, tt.inorder))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildTreeFromPreIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
