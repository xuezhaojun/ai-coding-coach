package pacific_atlantic

import (
	"sort"
	"testing"
)

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		name    string
		heights [][]int
		want    [][]int
	}{
		{
			name: "standard grid",
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			want: [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}},
		},
		{
			name:    "single cell",
			heights: [][]int{{1}},
			want:    [][]int{{0, 0}},
		},
		{
			name: "flat grid",
			heights: [][]int{
				{1, 1},
				{1, 1},
			},
			want: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
		},
		{
			name: "descending to corner",
			heights: [][]int{
				{3, 2},
				{2, 1},
			},
			want: [][]int{{0, 0}, {0, 1}, {1, 0}},
		},
		{
			name: "single row",
			heights: [][]int{
				{1, 2, 3},
			},
			want: [][]int{{0, 0}, {0, 1}, {0, 2}},
		},
		{
			name: "single column",
			heights: [][]int{
				{3},
				{2},
				{1},
			},
			want: [][]int{{0, 0}, {1, 0}, {2, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pacificAtlantic(tt.heights)
			// sort both for comparison
			sortCoords := func(coords [][]int) {
				sort.Slice(coords, func(i, j int) bool {
					if coords[i][0] != coords[j][0] {
						return coords[i][0] < coords[j][0]
					}
					return coords[i][1] < coords[j][1]
				})
			}
			sortCoords(got)
			sortCoords(tt.want)
			if len(got) != len(tt.want) {
				t.Errorf("pacificAtlantic() returned %d coords, want %d\ngot:  %v\nwant: %v", len(got), len(tt.want), got, tt.want)
				return
			}
			for i := range got {
				if got[i][0] != tt.want[i][0] || got[i][1] != tt.want[i][1] {
					t.Errorf("pacificAtlantic() mismatch at index %d: got %v, want %v", i, got[i], tt.want[i])
				}
			}
		})
	}
}
