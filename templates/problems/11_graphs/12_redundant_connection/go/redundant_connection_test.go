package redundant_connection

import (
	"reflect"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		name  string
		edges [][]int
		want  []int
	}{
		{
			name:  "triangle",
			edges: [][]int{{1, 2}, {1, 3}, {2, 3}},
			want:  []int{2, 3},
		},
		{
			name:  "four nodes with cycle",
			edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}},
			want:  []int{1, 4},
		},
		{
			name:  "last edge creates cycle",
			edges: [][]int{{1, 2}, {1, 3}, {3, 4}, {2, 4}},
			want:  []int{2, 4},
		},
		{
			name:  "two nodes",
			edges: [][]int{{1, 2}, {2, 1}},
			want:  []int{2, 1},
		},
		{
			name:  "five node cycle",
			edges: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 1}},
			want:  []int{5, 1},
		},
		{
			name:  "star with extra edge",
			edges: [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}},
			want:  []int{3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findRedundantConnection(tt.edges)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRedundantConnection(%v) = %v, want %v", tt.edges, got, tt.want)
			}
		})
	}
}
