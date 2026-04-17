package insert_interval

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		name        string
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{"merge in middle", [][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{"merge multiple", [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{"empty intervals", [][]int{}, []int{5, 7}, [][]int{{5, 7}}},
		{"insert at beginning", [][]int{{3, 5}, {6, 9}}, []int{1, 2}, [][]int{{1, 2}, {3, 5}, {6, 9}}},
		{"insert at end", [][]int{{1, 3}, {6, 9}}, []int{10, 12}, [][]int{{1, 3}, {6, 9}, {10, 12}}},
		{"merge all", [][]int{{1, 3}, {4, 6}, {7, 9}}, []int{0, 10}, [][]int{{0, 10}}},
		{"no overlap", [][]int{{1, 2}, {5, 6}}, []int{3, 4}, [][]int{{1, 2}, {3, 4}, {5, 6}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insert(tt.intervals, tt.newInterval)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("insert(%v, %v) = %v, want %v", tt.intervals, tt.newInterval, got, tt.expected)
			}
		})
	}
}
