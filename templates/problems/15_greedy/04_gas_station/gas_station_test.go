package gas_station

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		name     string
		gas      []int
		cost     []int
		expected int
	}{
		{"example 1", []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{"no solution", []int{2, 3, 4}, []int{3, 4, 3}, -1},
		{"single station enough", []int{5}, []int{3}, 0},
		{"single station not enough", []int{2}, []int{4}, -1},
		{"start at index 0", []int{3, 1, 1}, []int{1, 2, 2}, 0},
		{"all equal", []int{1, 1, 1}, []int{1, 1, 1}, 0},
		{"start at last", []int{1, 1, 5}, []int{2, 3, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canCompleteCircuit(tt.gas, tt.cost)
			if got != tt.expected {
				t.Errorf("canCompleteCircuit(%v, %v) = %d, want %d", tt.gas, tt.cost, got, tt.expected)
			}
		})
	}
}
