package remove_nth_from_end

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		n    int
		want []int
	}{
		{"remove last", []int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4}},
		{"remove first", []int{1, 2, 3, 4, 5}, 5, []int{2, 3, 4, 5}},
		{"remove middle", []int{1, 2, 3, 4, 5}, 3, []int{1, 2, 4, 5}},
		{"single element", []int{1}, 1, nil},
		{"two elements remove last", []int{1, 2}, 1, []int{1}},
		{"two elements remove first", []int{1, 2}, 2, []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.vals)
			got := listToSlice(removeNthFromEnd(head, tt.n))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
