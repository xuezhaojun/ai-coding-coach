package graph_valid_tree

import "testing"

func TestValidTree(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  bool
	}{
		{
			name:  "valid tree",
			n:     5,
			edges: [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			want:  true,
		},
		{
			name:  "has cycle",
			n:     5,
			edges: [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			want:  false,
		},
		{
			name:  "disconnected",
			n:     4,
			edges: [][]int{{0, 1}, {2, 3}},
			want:  false,
		},
		{
			name:  "single node",
			n:     1,
			edges: [][]int{},
			want:  true,
		},
		{
			name:  "two nodes connected",
			n:     2,
			edges: [][]int{{0, 1}},
			want:  true,
		},
		{
			name:  "too many edges",
			n:     3,
			edges: [][]int{{0, 1}, {1, 2}, {0, 2}},
			want:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := validTree(tt.n, tt.edges)
			if got != tt.want {
				t.Errorf("validTree(%d, %v) = %v, want %v", tt.n, tt.edges, got, tt.want)
			}
		})
	}
}
