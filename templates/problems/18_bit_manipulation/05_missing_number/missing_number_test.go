package missing_number

import "testing"

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"missing 2", []int{3, 0, 1}, 2},
		{"missing 2 from 3", []int{0, 1}, 2},
		{"missing 8", []int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{"missing 0", []int{1}, 0},
		{"missing 1", []int{0}, 1},
		{"single zero", []int{0, 2, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := missingNumber(tt.nums)
			if got != tt.expected {
				t.Errorf("missingNumber(%v) = %d, want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
