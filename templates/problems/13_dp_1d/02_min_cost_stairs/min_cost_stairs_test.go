package min_cost_stairs

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{
			name: "example 1",
			cost: []int{10, 15, 20},
			want: 15,
		},
		{
			name: "example 2",
			cost: []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
			want: 6,
		},
		{
			name: "two steps equal cost",
			cost: []int{5, 5},
			want: 5,
		},
		{
			name: "two steps different cost",
			cost: []int{1, 100},
			want: 1,
		},
		{
			name: "increasing cost",
			cost: []int{1, 2, 3, 4, 5},
			want: 6,
		},
		{
			name: "all zeros",
			cost: []int{0, 0, 0, 0},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCostClimbingStairs(tt.cost)
			if got != tt.want {
				t.Errorf("minCostClimbingStairs(%v) = %d, want %d", tt.cost, got, tt.want)
			}
		})
	}
}
