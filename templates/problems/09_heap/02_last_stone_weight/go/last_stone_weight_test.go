package last_stone_weight

import "testing"

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{
			name:   "example from problem",
			stones: []int{2, 7, 4, 1, 8, 1},
			want:   1,
		},
		{
			name:   "single stone",
			stones: []int{1},
			want:   1,
		},
		{
			name:   "two equal stones",
			stones: []int{3, 3},
			want:   0,
		},
		{
			name:   "two different stones",
			stones: []int{1, 5},
			want:   4,
		},
		{
			name:   "all same weight",
			stones: []int{2, 2, 2, 2},
			want:   0,
		},
		{
			name:   "descending weights",
			stones: []int{10, 5, 3, 1},
			want:   1,
		},
		{
			name:   "three stones",
			stones: []int{3, 7, 2},
			want:   2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lastStoneWeight(tt.stones)
			if got != tt.want {
				t.Errorf("lastStoneWeight(%v) = %d, want %d", tt.stones, got, tt.want)
			}
		})
	}
}
