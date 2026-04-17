package sliding_window_max

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "basic case",
			nums:     []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:        3,
			expected: []int{3, 3, 5, 5, 6, 7},
		},
		{
			name:     "k equals array length",
			nums:     []int{1, 3, 2},
			k:        3,
			expected: []int{3},
		},
		{
			name:     "k equals 1",
			nums:     []int{1, -1, 3},
			k:        1,
			expected: []int{1, -1, 3},
		},
		{
			name:     "single element",
			nums:     []int{5},
			k:        1,
			expected: []int{5},
		},
		{
			name:     "decreasing sequence",
			nums:     []int{7, 6, 5, 4, 3},
			k:        2,
			expected: []int{7, 6, 5, 4},
		},
		{
			name:     "increasing sequence",
			nums:     []int{1, 2, 3, 4, 5},
			k:        2,
			expected: []int{2, 3, 4, 5},
		},
		{
			name:     "all same values",
			nums:     []int{4, 4, 4, 4},
			k:        2,
			expected: []int{4, 4, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSlidingWindow(tt.nums, tt.k)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("maxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.expected)
			}
		})
	}
}
