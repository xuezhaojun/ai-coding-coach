package top_k_frequent

import (
	"reflect"
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{
			name: "basic case",
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			name: "single element",
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
		{
			name: "all same frequency k equals length",
			nums: []int{1, 2, 3},
			k:    3,
			want: []int{1, 2, 3},
		},
		{
			name: "negative numbers",
			nums: []int{-1, -1, -2, -2, -2, -3},
			k:    1,
			want: []int{-2},
		},
		{
			name: "k equals 1 with clear winner",
			nums: []int{4, 4, 4, 1, 2, 3},
			k:    1,
			want: []int{4},
		},
		{
			name: "two elements",
			nums: []int{1, 2},
			k:    2,
			want: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := topKFrequent(tt.nums, tt.k)
			if got == nil {
				t.Errorf("topKFrequent(%v, %d) = nil, want %v", tt.nums, tt.k, tt.want)
				return
			}
			// Sort for order-independent comparison
			sortedGot := make([]int, len(got))
			copy(sortedGot, got)
			sort.Ints(sortedGot)

			sortedWant := make([]int, len(tt.want))
			copy(sortedWant, tt.want)
			sort.Ints(sortedWant)

			if !reflect.DeepEqual(sortedGot, sortedWant) {
				t.Errorf("topKFrequent(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
