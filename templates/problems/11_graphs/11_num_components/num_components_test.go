package num_components

import "testing"

func TestCountComponents(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{
			name:  "two components",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {3, 4}},
			want:  2,
		},
		{
			name:  "one component",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}},
			want:  1,
		},
		{
			name:  "all isolated",
			n:     4,
			edges: [][]int{},
			want:  4,
		},
		{
			name:  "single node",
			n:     1,
			edges: [][]int{},
			want:  1,
		},
		{
			name:  "three components",
			n:     6,
			edges: [][]int{{0, 1}, {2, 3}, {4, 5}},
			want:  3,
		},
		{
			name:  "fully connected",
			n:     3,
			edges: [][]int{{0, 1}, {1, 2}, {0, 2}},
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countComponents(tt.n, tt.edges)
			if got != tt.want {
				t.Errorf("countComponents(%d, %v) = %v, want %v", tt.n, tt.edges, got, tt.want)
			}
		})
	}
}
