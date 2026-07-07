package level_order

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want [][]int
	}{
		{"example 1", []int{3, 9, 20, -101, -101, 15, 7}, [][]int{{3}, {9, 20}, {15, 7}}},
		{"example 2", []int{1}, [][]int{{1}}},
		{"empty tree", []int{}, nil},
		{"two levels", []int{1, 2, 3}, [][]int{{1}, {2, 3}}},
		{"left only", []int{1, 2, -101, 3}, [][]int{{1}, {2}, {3}}},
		{"full tree", []int{1, 2, 3, 4, 5, 6, 7}, [][]int{{1}, {2, 3}, {4, 5, 6, 7}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeHelper(tt.vals)
			got := levelOrder(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
