package stock_cooldown

import "testing"

func TestMaxProfitCooldown(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "example 1",
			prices: []int{1, 2, 3, 0, 2},
			want:   3,
		},
		{
			name:   "single day",
			prices: []int{1},
			want:   0,
		},
		{
			name:   "decreasing prices",
			prices: []int{5, 4, 3, 2, 1},
			want:   0,
		},
		{
			name:   "two days profit",
			prices: []int{1, 2},
			want:   1,
		},
		{
			name:   "two days no profit",
			prices: []int{2, 1},
			want:   0,
		},
		{
			name:   "alternating",
			prices: []int{1, 4, 2, 7},
			want:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfitCooldown(tt.prices)
			if got != tt.want {
				t.Errorf("maxProfitCooldown(%v) = %d, want %d", tt.prices, got, tt.want)
			}
		})
	}
}
