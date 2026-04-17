package merge_intervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		expected  [][]int
	}{
		{"overlapping", [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{"touching", [][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{"no overlap", [][]int{{1, 2}, {4, 5}}, [][]int{{1, 2}, {4, 5}}},
		{"single interval", [][]int{{1, 5}}, [][]int{{1, 5}}},
		{"all merge", [][]int{{1, 4}, {2, 5}, {3, 6}}, [][]int{{1, 6}}},
		{"unsorted input", [][]int{{3, 4}, {1, 2}, {5, 6}}, [][]int{{1, 2}, {3, 4}, {5, 6}}},
		{"contained interval", [][]int{{1, 10}, {3, 5}}, [][]int{{1, 10}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.intervals)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("merge(%v) = %v, want %v", tt.intervals, got, tt.expected)
			}
		})
	}
}
