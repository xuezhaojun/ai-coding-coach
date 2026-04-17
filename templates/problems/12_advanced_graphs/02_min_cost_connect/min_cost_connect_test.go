package min_cost_connect

import "testing"

func TestMinCostConnectPoints(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "five points",
			points: [][]int{{0, 0}, {2, 2}, {3, 10}, {5, 2}, {7, 0}},
			want:   20,
		},
		{
			name:   "three collinear points",
			points: [][]int{{3, 12}, {-2, 5}, {-4, 1}},
			want:   18,
		},
		{
			name:   "single point",
			points: [][]int{{0, 0}},
			want:   0,
		},
		{
			name:   "two points",
			points: [][]int{{0, 0}, {1, 1}},
			want:   2,
		},
		{
			name:   "square corners",
			points: [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}},
			want:   3,
		},
		{
			name:   "negative coordinates",
			points: [][]int{{-1, -1}, {-3, -3}, {1, 1}},
			want:   8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCostConnectPoints(tt.points)
			if got != tt.want {
				t.Errorf("minCostConnectPoints(%v) = %v, want %v", tt.points, got, tt.want)
			}
		})
	}
}
