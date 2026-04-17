package single_number

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1", []int{2, 2, 1}, 1},
		{"example 2", []int{4, 1, 2, 1, 2}, 4},
		{"single element", []int{1}, 1},
		{"negative numbers", []int{-1, -1, -2}, -2},
		{"zero is single", []int{0, 1, 1}, 0},
		{"larger array", []int{1, 3, 1, 2, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := singleNumber(tt.nums)
			if got != tt.expected {
				t.Errorf("singleNumber(%v) = %d, want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
