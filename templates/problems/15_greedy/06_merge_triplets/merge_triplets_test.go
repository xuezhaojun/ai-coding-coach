package merge_triplets

import "testing"

func TestMergeTriplets(t *testing.T) {
	tests := []struct {
		name     string
		triplets [][]int
		target   []int
		expected bool
	}{
		{"example true", [][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, []int{2, 7, 5}, true},
		{"example false", [][]int{{3, 4, 5}, {4, 5, 6}}, []int{3, 2, 5}, false},
		{"exact match single", [][]int{{2, 5, 3}}, []int{2, 5, 3}, true},
		{"no valid triplet", [][]int{{1, 1, 1}}, []int{2, 2, 2}, false},
		{"filter out exceeding", [][]int{{2, 5, 3}, {10, 1, 1}, {1, 7, 5}}, []int{2, 7, 5}, true},
		{"all exceed one dim", [][]int{{3, 1, 1}, {3, 2, 2}}, []int{2, 2, 2}, false},
		{"multiple combos", [][]int{{1, 2, 3}, {2, 1, 3}, {2, 2, 1}}, []int{2, 2, 3}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTriplets(tt.triplets, tt.target)
			if got != tt.expected {
				t.Errorf("mergeTriplets(%v, %v) = %v, want %v", tt.triplets, tt.target, got, tt.expected)
			}
		})
	}
}
