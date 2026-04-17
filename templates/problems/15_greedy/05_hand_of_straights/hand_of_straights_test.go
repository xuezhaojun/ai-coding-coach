package hand_of_straights

import "testing"

func TestIsNStraightHand(t *testing.T) {
	tests := []struct {
		name      string
		hand      []int
		groupSize int
		expected  bool
	}{
		{"example true", []int{1, 2, 3, 6, 2, 3, 4, 7, 8}, 3, true},
		{"example false", []int{1, 2, 3, 4, 5}, 4, false},
		{"single group", []int{1, 2, 3}, 3, true},
		{"group size 1", []int{5, 3, 1}, 1, true},
		{"not divisible", []int{1, 2, 3, 4}, 3, false},
		{"duplicates needed", []int{1, 1, 2, 2, 3, 3}, 3, true},
		{"gap in sequence", []int{1, 2, 4, 5, 6, 7}, 3, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isNStraightHand(tt.hand, tt.groupSize)
			if got != tt.expected {
				t.Errorf("isNStraightHand(%v, %d) = %v, want %v", tt.hand, tt.groupSize, got, tt.expected)
			}
		})
	}
}
