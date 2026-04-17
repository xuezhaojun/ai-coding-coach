package walls_and_gates

import (
	"reflect"
	"testing"
)

const INF = 2147483647

func TestWallsAndGates(t *testing.T) {
	tests := []struct {
		name  string
		rooms [][]int
		want  [][]int
	}{
		{
			name: "standard grid",
			rooms: [][]int{
				{INF, -1, 0, INF},
				{INF, INF, INF, -1},
				{INF, -1, INF, -1},
				{0, -1, INF, INF},
			},
			want: [][]int{
				{3, -1, 0, 1},
				{2, 2, 1, -1},
				{1, -1, 2, -1},
				{0, -1, 3, 4},
			},
		},
		{
			name:  "single gate",
			rooms: [][]int{{0}},
			want:  [][]int{{0}},
		},
		{
			name:  "single wall",
			rooms: [][]int{{-1}},
			want:  [][]int{{-1}},
		},
		{
			name:  "no gates",
			rooms: [][]int{{INF, INF}, {INF, INF}},
			want:  [][]int{{INF, INF}, {INF, INF}},
		},
		{
			name: "all gates",
			rooms: [][]int{
				{0, 0},
				{0, 0},
			},
			want: [][]int{
				{0, 0},
				{0, 0},
			},
		},
		{
			name: "gate in corner",
			rooms: [][]int{
				{0, INF},
				{INF, INF},
			},
			want: [][]int{
				{0, 1},
				{1, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wallsAndGates(tt.rooms)
			if !reflect.DeepEqual(tt.rooms, tt.want) {
				t.Errorf("wallsAndGates() = %v, want %v", tt.rooms, tt.want)
			}
		})
	}
}
