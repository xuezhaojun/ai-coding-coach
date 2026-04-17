package find_min_rotated

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"rotated", []int{3, 4, 5, 1, 2}, 1},
		{"rotated once", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"not rotated", []int{11, 13, 15, 17}, 11},
		{"single element", []int{1}, 1},
		{"two elements rotated", []int{2, 1}, 1},
		{"two elements sorted", []int{1, 2}, 1},
		{"min at end", []int{2, 3, 4, 5, 1}, 1},
		{"min at start", []int{1, 2, 3, 4, 5}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.nums); got != tt.want {
				t.Errorf("findMin(%v) = %d, want %d", tt.nums, got, tt.want)
			}
		})
	}
}
