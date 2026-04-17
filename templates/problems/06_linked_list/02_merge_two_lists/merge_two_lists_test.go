package merge_two_lists

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name  string
		list1 []int
		list2 []int
		want  []int
	}{
		{"both nil", nil, nil, nil},
		{"first nil", nil, []int{1, 2, 3}, []int{1, 2, 3}},
		{"second nil", []int{1, 2, 3}, nil, []int{1, 2, 3}},
		{"interleaved", []int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"one before other", []int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"duplicates", []int{1, 1, 3}, []int{1, 2, 4}, []int{1, 1, 1, 2, 3, 4}},
		{"single elements", []int{5}, []int{1}, []int{1, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l1 := buildList(tt.list1)
			l2 := buildList(tt.list2)
			got := listToSlice(mergeTwoLists(l1, l2))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
