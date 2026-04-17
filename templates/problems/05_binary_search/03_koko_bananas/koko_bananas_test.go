package koko_bananas

import "testing"

func TestMinEatingSpeed(t *testing.T) {
	tests := []struct {
		name  string
		piles []int
		h     int
		want  int
	}{
		{"example 1", []int{3, 6, 7, 11}, 8, 4},
		{"example 2", []int{30, 11, 23, 4, 20}, 5, 30},
		{"example 3", []int{30, 11, 23, 4, 20}, 6, 23},
		{"single pile exact", []int{10}, 1, 10},
		{"single pile slow", []int{10}, 10, 1},
		{"all ones", []int{1, 1, 1}, 3, 1},
		{"large h", []int{3, 6, 7, 11}, 100, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEatingSpeed(tt.piles, tt.h); got != tt.want {
				t.Errorf("minEatingSpeed(%v, %d) = %d, want %d", tt.piles, tt.h, got, tt.want)
			}
		})
	}
}
