package min_interval_query

import (
	"reflect"
	"testing"
)

func TestMinInterval(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		queries   []int
		expected  []int
	}{
		{"example 1", [][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}}, []int{2, 3, 4, 5}, []int{3, 3, 1, 4}},
		{"example 2", [][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}}, []int{2, 19, 5, 22}, []int{2, -1, 4, 6}},
		{"no intervals", [][]int{}, []int{1, 2}, []int{-1, -1}},
		{"query outside all", [][]int{{1, 3}}, []int{5}, []int{-1}},
		{"single point interval", [][]int{{5, 5}}, []int{5}, []int{1}},
		{"multiple covering", [][]int{{1, 10}, {2, 5}, {3, 4}}, []int{3}, []int{2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minInterval(tt.intervals, tt.queries)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("minInterval(%v, %v) = %v, want %v", tt.intervals, tt.queries, got, tt.expected)
			}
		})
	}
}
