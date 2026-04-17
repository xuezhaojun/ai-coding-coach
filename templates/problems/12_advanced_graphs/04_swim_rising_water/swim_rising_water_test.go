package swim_rising_water

import "testing"

func TestSwimInWater(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{
			name: "2x2 grid",
			grid: [][]int{
				{0, 2},
				{1, 3},
			},
			want: 3,
		},
		{
			name: "5x5 grid",
			grid: [][]int{
				{0, 1, 2, 3, 4},
				{24, 23, 22, 21, 5},
				{12, 13, 14, 15, 16},
				{11, 17, 18, 19, 20},
				{10, 9, 8, 7, 6},
			},
			want: 16,
		},
		{
			name: "single cell",
			grid: [][]int{{0}},
			want: 0,
		},
		{
			name: "3x3 grid",
			grid: [][]int{
				{0, 1, 2},
				{5, 4, 3},
				{6, 7, 8},
			},
			want: 8,
		},
		{
			name: "direct path best",
			grid: [][]int{
				{0, 3, 5},
				{1, 4, 6},
				{2, 7, 8},
			},
			want: 8,
		},
		{
			name: "2x2 easy",
			grid: [][]int{
				{0, 1},
				{2, 3},
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := swimInWater(tt.grid)
			if got != tt.want {
				t.Errorf("swimInWater(%v) = %v, want %v", tt.grid, got, tt.want)
			}
		})
	}
}
