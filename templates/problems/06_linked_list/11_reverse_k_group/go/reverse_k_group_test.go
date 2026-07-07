package reverse_k_group

import (
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		k    int
		want []int
	}{
		{"nil list", nil, 2, nil},
		{"k equals 1", []int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
		{"k equals 2", []int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{"k equals 3", []int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{"exact groups", []int{1, 2, 3, 4}, 2, []int{2, 1, 4, 3}},
		{"k equals length", []int{1, 2, 3}, 3, []int{3, 2, 1}},
		{"k greater than length", []int{1, 2}, 3, []int{1, 2}},
		{"single element", []int{1}, 1, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.vals)
			got := listToSlice(reverseKGroup(head, tt.k))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
