package best_time_stock

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "basic profit",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "no profit possible",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "single day",
			prices:   []int{5},
			expected: 0,
		},
		{
			name:     "two days profit",
			prices:   []int{1, 2},
			expected: 1,
		},
		{
			name:     "two days no profit",
			prices:   []int{2, 1},
			expected: 0,
		},
		{
			name:     "buy at start sell at end",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "all same price",
			prices:   []int{3, 3, 3, 3},
			expected: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.expected {
				t.Errorf("maxProfit(%v) = %d, want %d", tt.prices, got, tt.expected)
			}
		})
	}
}
