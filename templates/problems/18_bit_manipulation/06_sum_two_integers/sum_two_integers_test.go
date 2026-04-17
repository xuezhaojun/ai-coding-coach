package sum_two_integers

import "testing"

func TestGetSum(t *testing.T) {
	tests := []struct {
		name     string
		a        int
		b        int
		expected int
	}{
		{"both positive", 1, 2, 3},
		{"positive and negative", 2, -1, 1},
		{"both negative", -1, -1, -2},
		{"zero and number", 0, 5, 5},
		{"both zero", 0, 0, 0},
		{"larger numbers", 100, 200, 300},
		{"negative result", -5, 3, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSum(tt.a, tt.b)
			if got != tt.expected {
				t.Errorf("getSum(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		})
	}
}
