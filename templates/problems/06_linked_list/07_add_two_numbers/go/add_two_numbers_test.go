package add_two_numbers

import (
	"reflect"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{"both zero", []int{0}, []int{0}, []int{0}},
		{"no carry", []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{"with carry", []int{9, 9, 9}, []int{1}, []int{0, 0, 0, 1}},
		{"different lengths", []int{1, 8}, []int{0}, []int{1, 8}},
		{"single digits", []int{5}, []int{5}, []int{0, 1}},
		{"large carry chain", []int{9, 9, 9, 9}, []int{1}, []int{0, 0, 0, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := buildList(tt.l1)
			l2 := buildList(tt.l2)
			got := listToSlice(addTwoNumbers(l1, l2))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
