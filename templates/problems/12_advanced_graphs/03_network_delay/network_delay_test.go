package network_delay

import "testing"

func TestNetworkDelayTime(t *testing.T) {
	tests := []struct {
		name  string
		times [][]int
		n     int
		k     int
		want  int
	}{
		{
			name:  "standard case",
			times: [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}},
			n:     4,
			k:     2,
			want:  2,
		},
		{
			name:  "unreachable node",
			times: [][]int{{1, 2, 1}},
			n:     2,
			k:     2,
			want:  -1,
		},
		{
			name:  "single node",
			times: [][]int{},
			n:     1,
			k:     1,
			want:  0,
		},
		{
			name:  "two paths different weights",
			times: [][]int{{1, 2, 1}, {1, 3, 4}, {2, 3, 2}},
			n:     3,
			k:     1,
			want:  3,
		},
		{
			name:  "all connected from source",
			times: [][]int{{1, 2, 5}, {1, 3, 2}, {1, 4, 9}},
			n:     4,
			k:     1,
			want:  9,
		},
		{
			name:  "chain of nodes",
			times: [][]int{{1, 2, 1}, {2, 3, 2}, {3, 4, 3}},
			n:     4,
			k:     1,
			want:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := networkDelayTime(tt.times, tt.n, tt.k)
			if got != tt.want {
				t.Errorf("networkDelayTime(%v, %d, %d) = %v, want %v", tt.times, tt.n, tt.k, got, tt.want)
			}
		})
	}
}
