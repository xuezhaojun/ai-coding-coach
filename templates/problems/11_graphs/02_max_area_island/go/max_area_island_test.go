package max_area_island

import "testing"

func TestMaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "mixed islands",
			grid: [][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			want: 6,
		},
		{
			name: "all water",
			grid: [][]int{
				{0, 0, 0},
				{0, 0, 0},
			},
			want: 0,
		},
		{
			name: "single cell island",
			grid: [][]int{
				{0, 1, 0},
				{0, 0, 0},
			},
			want: 1,
		},
		{
			name: "entire grid is one island",
			grid: [][]int{
				{1, 1},
				{1, 1},
			},
			want: 4,
		},
		{
			name: "L-shaped island",
			grid: [][]int{
				{1, 0},
				{1, 0},
				{1, 1},
			},
			want: 4,
		},
		{
			name: "two separate islands",
			grid: [][]int{
				{1, 0, 1, 1},
				{0, 0, 1, 1},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// deep copy grid
			g := make([][]int, len(tt.grid))
			for i := range tt.grid {
				g[i] = make([]int, len(tt.grid[i]))
				copy(g[i], tt.grid[i])
			}
			got := maxAreaOfIsland(g)
			if got != tt.want {
				t.Errorf("maxAreaOfIsland() = %v, want %v", got, tt.want)
			}
		})
	}
}
