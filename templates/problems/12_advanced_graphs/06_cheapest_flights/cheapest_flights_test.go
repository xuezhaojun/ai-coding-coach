package cheapest_flights

import "testing"

func TestFindCheapestPrice(t *testing.T) {
	tests := []struct {
		name    string
		n       int
		flights [][]int
		src     int
		dst     int
		k       int
		want    int
	}{
		{
			name:    "standard case",
			n:       4,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}},
			src:     0,
			dst:     3,
			k:       1,
			want:    700,
		},
		{
			name:    "direct flight",
			n:       3,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			src:     0,
			dst:     2,
			k:       0,
			want:    500,
		},
		{
			name:    "cheaper with stop",
			n:       3,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}},
			src:     0,
			dst:     2,
			k:       1,
			want:    200,
		},
		{
			name:    "no route",
			n:       3,
			flights: [][]int{{0, 1, 100}},
			src:     0,
			dst:     2,
			k:       1,
			want:    -1,
		},
		{
			name:    "not enough stops",
			n:       4,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {2, 3, 100}},
			src:     0,
			dst:     3,
			k:       1,
			want:    -1,
		},
		{
			name:    "same src and dst",
			n:       2,
			flights: [][]int{{0, 1, 100}},
			src:     0,
			dst:     0,
			k:       0,
			want:    0,
		},
		{
			name:    "two stops allowed",
			n:       4,
			flights: [][]int{{0, 1, 100}, {1, 2, 100}, {2, 3, 100}},
			src:     0,
			dst:     3,
			k:       2,
			want:    300,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findCheapestPrice(tt.n, tt.flights, tt.src, tt.dst, tt.k)
			if got != tt.want {
				t.Errorf("findCheapestPrice(%d, %v, %d, %d, %d) = %v, want %v",
					tt.n, tt.flights, tt.src, tt.dst, tt.k, got, tt.want)
			}
		})
	}
}
