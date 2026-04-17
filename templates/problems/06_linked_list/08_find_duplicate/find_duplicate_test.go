package find_duplicate

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"simple", []int{1, 3, 4, 2, 2}, 2},
		{"duplicate three", []int{3, 1, 3, 4, 2}, 3},
		{"all same", []int{1, 1, 1, 1, 1}, 1},
		{"two elements", []int{1, 1}, 1},
		{"duplicate at end", []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1}, 9},
		{"duplicate two", []int{3, 3, 3, 3, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
