package max_subarray

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"mixed positive and negative", []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{"single element positive", []int{1}, 1},
		{"single element negative", []int{-1}, -1},
		{"all negative", []int{-2, -3, -1, -5}, -1},
		{"all positive", []int{1, 2, 3, 4}, 10},
		{"two elements", []int{-1, 2}, 2},
		{"negative then positive", []int{-2, -1, 3, 4, -1, 2}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxSubArray(tt.nums)
			if got != tt.expected {
				t.Errorf("maxSubArray(%v) = %d, want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
