package right_side_view

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"example 1", []int{1, 2, 3, -101, 5, -101, 4}, []int{1, 3, 4}},
		{"example 2", []int{1, -101, 3}, []int{1, 3}},
		{"empty tree", []int{}, nil},
		{"single node", []int{1}, []int{1}},
		{"left deeper", []int{1, 2, 3, 4}, []int{1, 3, 4}},
		{"all left", []int{1, 2, -101, 3}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeHelper(tt.vals)
			got := rightSideView(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
