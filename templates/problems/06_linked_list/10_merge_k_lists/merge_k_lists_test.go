package merge_k_lists

import (
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		name  string
		lists [][]int
		want  []int
	}{
		{"nil input", nil, nil},
		{"empty lists", [][]int{{}, {}, {}}, nil},
		{"single list", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"two lists", [][]int{{1, 4, 5}, {1, 3, 4}}, []int{1, 1, 3, 4, 4, 5}},
		{"three lists", [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, []int{1, 1, 2, 3, 4, 4, 5, 6}},
		{"with empty list", [][]int{{1, 2}, {}, {3, 4}}, []int{1, 2, 3, 4}},
		{"all single elements", [][]int{{5}, {1}, {3}}, []int{1, 3, 5}},
		{"duplicates across lists", [][]int{{1, 1}, {1, 1}}, []int{1, 1, 1, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var lists []*ListNode
			for _, vals := range tt.lists {
				lists = append(lists, buildList(vals))
			}
			got := listToSlice(mergeKLists(lists))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
