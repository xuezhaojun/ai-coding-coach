package rotting_oranges

import "testing"

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "standard case",
			grid: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			want: 4,
		},
		{
			name: "impossible to rot all",
			grid: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			want: -1,
		},
		{
			name: "no fresh oranges",
			grid: [][]int{
				{0, 2},
			},
			want: 0,
		},
		{
			name: "empty grid cells only",
			grid: [][]int{
				{0},
			},
			want: 0,
		},
		{
			name: "already all rotten",
			grid: [][]int{
				{2, 2},
				{2, 2},
			},
			want: 0,
		},
		{
			name: "single fresh unreachable",
			grid: [][]int{
				{1},
			},
			want: -1,
		},
		{
			name: "one step",
			grid: [][]int{
				{2, 1},
			},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := orangesRotting(tt.grid)
			if got != tt.want {
				t.Errorf("orangesRotting() = %v, want %v", got, tt.want)
			}
		})
	}
}
