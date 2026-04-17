package binary_search

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"found middle", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"found first", []int{-1, 0, 3, 5, 9, 12}, -1, 0},
		{"found last", []int{-1, 0, 3, 5, 9, 12}, 12, 5},
		{"not found", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{"single element found", []int{5}, 5, 0},
		{"single element not found", []int{5}, 3, -1},
		{"two elements", []int{1, 3}, 3, 1},
		{"empty array", []int{}, 1, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search(%v, %d) = %d, want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
