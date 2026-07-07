package two_sum_ii

import (
	"reflect"
	"testing"
)

func TestTwoSumII(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		target   int
		expected []int
	}{
		{
			name:     "basic case",
			numbers:  []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			name:     "middle elements",
			numbers:  []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		{
			name:     "negative numbers",
			numbers:  []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
		{
			name:     "larger array",
			numbers:  []int{1, 2, 3, 4, 4, 9, 56, 90},
			target:   8,
			expected: []int{4, 5},
		},
		{
			name:     "first and last",
			numbers:  []int{1, 3, 5, 7},
			target:   8,
			expected: []int{1, 4},
		},
		{
			name:     "two elements",
			numbers:  []int{5, 25},
			target:   30,
			expected: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSumII(tt.numbers, tt.target)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("twoSumII(%v, %d) = %v, want %v", tt.numbers, tt.target, got, tt.expected)
			}
		})
	}
}
