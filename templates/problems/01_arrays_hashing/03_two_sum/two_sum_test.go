package two_sum

import (
	"sort"
	"testing"

	"reflect"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "basic case",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "elements not adjacent",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "same element value",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "negative numbers",
			nums:   []int{-1, -2, -3, -4, -5},
			target: -8,
			want:   []int{2, 4},
		},
		{
			name:   "mixed positive and negative",
			nums:   []int{-3, 4, 3, 90},
			target: 0,
			want:   []int{0, 2},
		},
		{
			name:   "large array target at end",
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			target: 19,
			want:   []int{8, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if got == nil {
				t.Errorf("twoSum(%v, %d) = nil, want %v", tt.nums, tt.target, tt.want)
				return
			}
			// Sort both slices for order-independent comparison
			sortedGot := make([]int, len(got))
			copy(sortedGot, got)
			sort.Ints(sortedGot)

			sortedWant := make([]int, len(tt.want))
			copy(sortedWant, tt.want)
			sort.Ints(sortedWant)

			if !reflect.DeepEqual(sortedGot, sortedWant) {
				t.Errorf("twoSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
