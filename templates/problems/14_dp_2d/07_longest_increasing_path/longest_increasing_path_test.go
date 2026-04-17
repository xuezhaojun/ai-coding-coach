package longest_increasing_path

import "testing"

func TestLongestIncreasingPath(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{
			name:   "example 1",
			matrix: [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}},
			want:   4,
		},
		{
			name:   "example 2",
			matrix: [][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}},
			want:   4,
		},
		{
			name:   "single cell",
			matrix: [][]int{{1}},
			want:   1,
		},
		{
			name:   "single row",
			matrix: [][]int{{1, 2, 3, 4}},
			want:   4,
		},
		{
			name:   "single column",
			matrix: [][]int{{1}, {2}, {3}},
			want:   3,
		},
		{
			name:   "all same values",
			matrix: [][]int{{5, 5}, {5, 5}},
			want:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestIncreasingPath(tt.matrix)
			if got != tt.want {
				t.Errorf("longestIncreasingPath(%v) = %d, want %d", tt.matrix, got, tt.want)
			}
		})
	}
}
