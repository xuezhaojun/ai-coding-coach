package reorder_list

import (
	"reflect"
	"testing"
)

func TestReorderList(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"nil list", nil, nil},
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{1, 2}},
		{"three elements", []int{1, 2, 3}, []int{1, 3, 2}},
		{"four elements", []int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{"five elements", []int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
		{"six elements", []int{1, 2, 3, 4, 5, 6}, []int{1, 6, 2, 5, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.vals)
			reorderList(head)
			got := listToSlice(head)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderList() = %v, want %v", got, tt.want)
			}
		})
	}
}
