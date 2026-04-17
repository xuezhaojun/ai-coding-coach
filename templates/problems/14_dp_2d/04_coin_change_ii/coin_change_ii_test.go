package coin_change_ii

import "testing"

func TestChange(t *testing.T) {
	tests := []struct {
		name   string
		amount int
		coins  []int
		want   int
	}{
		{
			name:   "example 1",
			amount: 5,
			coins:  []int{1, 2, 5},
			want:   4,
		},
		{
			name:   "example 2",
			amount: 3,
			coins:  []int{2},
			want:   0,
		},
		{
			name:   "zero amount",
			amount: 0,
			coins:  []int{1, 2},
			want:   1,
		},
		{
			name:   "single coin",
			amount: 10,
			coins:  []int{10},
			want:   1,
		},
		{
			name:   "single penny",
			amount: 5,
			coins:  []int{1},
			want:   1,
		},
		{
			name:   "two coins",
			amount: 4,
			coins:  []int{1, 2},
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := change(tt.amount, tt.coins)
			if got != tt.want {
				t.Errorf("change(%d, %v) = %d, want %d", tt.amount, tt.coins, got, tt.want)
			}
		})
	}
}
