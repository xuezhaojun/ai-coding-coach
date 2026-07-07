package num_islands

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		name string
		grid [][]byte
		want int
	}{
		{
			name: "single island",
			grid: [][]byte{
				{'1', '1', '1'},
				{'0', '1', '0'},
				{'1', '1', '1'},
			},
			want: 1,
		},
		{
			name: "three islands",
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			want: 3,
		},
		{
			name: "all water",
			grid: [][]byte{
				{'0', '0', '0'},
				{'0', '0', '0'},
			},
			want: 0,
		},
		{
			name: "all land",
			grid: [][]byte{
				{'1', '1'},
				{'1', '1'},
			},
			want: 1,
		},
		{
			name: "diagonal islands",
			grid: [][]byte{
				{'1', '0'},
				{'0', '1'},
			},
			want: 2,
		},
		{
			name: "single cell land",
			grid: [][]byte{
				{'1'},
			},
			want: 1,
		},
		{
			name: "single cell water",
			grid: [][]byte{
				{'0'},
			},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// deep copy grid since it may be modified
			g := make([][]byte, len(tt.grid))
			for i := range tt.grid {
				g[i] = make([]byte, len(tt.grid[i]))
				copy(g[i], tt.grid[i])
			}
			got := numIslands(g)
			if got != tt.want {
				t.Errorf("numIslands() = %v, want %v", got, tt.want)
			}
		})
	}
}
