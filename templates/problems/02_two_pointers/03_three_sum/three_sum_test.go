package three_sum

import (
	"reflect"
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected [][]int
	}{
		{
			name:     "basic case",
			nums:     []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name:     "no triplets",
			nums:     []int{0, 1, 1},
			expected: [][]int{},
		},
		{
			name:     "all zeros",
			nums:     []int{0, 0, 0},
			expected: [][]int{{0, 0, 0}},
		},
		{
			name:     "empty input",
			nums:     []int{},
			expected: [][]int{},
		},
		{
			name:     "two elements only",
			nums:     []int{-1, 1},
			expected: [][]int{},
		},
		{
			name:     "multiple triplets",
			nums:     []int{-2, 0, 1, 1, 2},
			expected: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
		{
			name:     "all positive",
			nums:     []int{1, 2, 3, 4, 5},
			expected: [][]int{},
		},
		{
			name:     "duplicate left right values",
			nums:     []int{-2, 0, 0, 2, 2},
			expected: [][]int{{-2, 0, 2}},
		},
		{
			name:     "all same value after anchor",
			nums:     []int{-4, 2, 2, 2, 2},
			expected: [][]int{{-4, 2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			// Sort each triplet and the overall result for consistent comparison
			for _, triplet := range got {
				sort.Ints(triplet)
			}
			sort.Slice(got, func(i, j int) bool {
				for k := 0; k < 3; k++ {
					if got[i][k] != got[j][k] {
						return got[i][k] < got[j][k]
					}
				}
				return false
			})
			if got == nil {
				got = [][]int{}
			}
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("threeSum(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
