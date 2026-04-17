package coin_change

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name   string
		coins  []int
		amount int
		want   int
	}{
		{
			name:   "example 1",
			coins:  []int{1, 2, 5},
			amount: 11,
			want:   3,
		},
		{
			name:   "impossible",
			coins:  []int{2},
			amount: 3,
			want:   -1,
		},
		{
			name:   "zero amount",
			coins:  []int{1},
			amount: 0,
			want:   0,
		},
		{
			name:   "single coin exact",
			coins:  []int{1},
			amount: 1,
			want:   1,
		},
		{
			name:   "single coin multiple",
			coins:  []int{1},
			amount: 5,
			want:   5,
		},
		{
			name:   "large coins small amount",
			coins:  []int{5, 10},
			amount: 3,
			want:   -1,
		},
		{
			name:   "multiple denominations",
			coins:  []int{1, 5, 10, 25},
			amount: 30,
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := coinChange(tt.coins, tt.amount)
			if got != tt.want {
				t.Errorf("coinChange(%v, %d) = %d, want %d", tt.coins, tt.amount, got, tt.want)
			}
		})
	}
}
