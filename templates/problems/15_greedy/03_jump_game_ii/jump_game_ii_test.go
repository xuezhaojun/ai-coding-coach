package jump_game_ii

import "testing"

func TestJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"example 1", []int{2, 3, 1, 1, 4}, 2},
		{"example 2", []int{2, 3, 0, 1, 4}, 2},
		{"single element", []int{0}, 0},
		{"two elements", []int{1, 2}, 1},
		{"already at end", []int{1}, 0},
		{"linear jumps", []int{1, 1, 1, 1}, 3},
		{"large first jump", []int{5, 1, 1, 1, 1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jump(tt.nums)
			if got != tt.expected {
				t.Errorf("jump(%v) = %d, want %d", tt.nums, got, tt.expected)
			}
		})
	}
}
