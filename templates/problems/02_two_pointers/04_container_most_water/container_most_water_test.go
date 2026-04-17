package container_most_water

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "basic case",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "two elements",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "decreasing heights",
			height:   []int{4, 3, 2, 1, 4},
			expected: 16,
		},
		{
			name:     "equal heights",
			height:   []int{5, 5, 5, 5},
			expected: 15,
		},
		{
			name:     "one tall wall",
			height:   []int{1, 2, 1},
			expected: 2,
		},
		{
			name:     "large difference",
			height:   []int{1, 1000, 1000, 1},
			expected: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.expected {
				t.Errorf("maxArea(%v) = %d, want %d", tt.height, got, tt.expected)
			}
		})
	}
}
