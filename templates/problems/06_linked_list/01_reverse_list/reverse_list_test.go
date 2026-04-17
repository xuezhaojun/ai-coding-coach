package reverse_list

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"nil list", nil, nil},
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{1, 2}, []int{2, 1}},
		{"multiple elements", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"duplicates", []int{1, 1, 2, 2}, []int{2, 2, 1, 1}},
		{"negative values", []int{-1, 0, 1}, []int{1, 0, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.vals)
			got := listToSlice(reverseList(head))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
