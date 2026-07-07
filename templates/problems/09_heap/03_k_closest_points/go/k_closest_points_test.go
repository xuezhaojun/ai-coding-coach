package k_closest_points

import (
	"sort"
	"testing"
)

func TestKClosest(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		k      int
		want   [][]int
	}{
		{
			name:   "example 1",
			points: [][]int{{1, 3}, {-2, 2}},
			k:      1,
			want:   [][]int{{-2, 2}},
		},
		{
			name:   "example 2",
			points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:      2,
			want:   [][]int{{-2, 4}, {3, 3}},
		},
		{
			name:   "single point",
			points: [][]int{{0, 1}},
			k:      1,
			want:   [][]int{{0, 1}},
		},
		{
			name:   "all points same distance",
			points: [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}},
			k:      2,
			want:   [][]int{{-1, 0}, {0, -1}}, // any 2 of the 4
		},
		{
			name:   "origin point",
			points: [][]int{{0, 0}, {1, 1}, {2, 2}},
			k:      1,
			want:   [][]int{{0, 0}},
		},
		{
			name:   "k equals length",
			points: [][]int{{1, 2}, {3, 4}},
			k:      2,
			want:   [][]int{{1, 2}, {3, 4}},
		},
	}

	sortPoints := func(pts [][]int) {
		sort.Slice(pts, func(i, j int) bool {
			if pts[i][0] != pts[j][0] {
				return pts[i][0] < pts[j][0]
			}
			return pts[i][1] < pts[j][1]
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kClosest(tt.points, tt.k)
			if len(got) != len(tt.want) {
				t.Fatalf("kClosest() returned %d points, want %d", len(got), len(tt.want))
			}
			sortPoints(got)
			sortPoints(tt.want)
			// For the "all same distance" case, just check the count
			if tt.name == "all points same distance" {
				if len(got) != tt.k {
					t.Errorf("expected %d points, got %d", tt.k, len(got))
				}
				return
			}
			for i := range got {
				if got[i][0] != tt.want[i][0] || got[i][1] != tt.want[i][1] {
					t.Errorf("kClosest() = %v, want %v", got, tt.want)
					break
				}
			}
		})
	}
}
