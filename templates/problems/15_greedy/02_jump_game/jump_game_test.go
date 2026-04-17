package jump_game

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected bool
	}{
		{"reachable", []int{2, 3, 1, 1, 4}, true},
		{"not reachable", []int{3, 2, 1, 0, 4}, false},
		{"single element", []int{0}, true},
		{"two elements reachable", []int{1, 0}, true},
		{"all zeros except first", []int{5, 0, 0, 0, 0}, true},
		{"large first jump", []int{10, 0, 0, 0, 0, 0, 0, 0, 1}, true},
		{"stuck at zero", []int{1, 0, 1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canJump(tt.nums)
			if got != tt.expected {
				t.Errorf("canJump(%v) = %v, want %v", tt.nums, got, tt.expected)
			}
		})
	}
}
