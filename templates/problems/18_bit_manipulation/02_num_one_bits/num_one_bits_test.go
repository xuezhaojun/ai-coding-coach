package num_one_bits

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		name     string
		n        uint32
		expected int
	}{
		{"three ones", 0b00000000000000000000000000001011, 3},
		{"one bit", 0b00000000000000000000000010000000, 1},
		{"all ones 32bit", 0b11111111111111111111111111111101, 31},
		{"zero", 0, 0},
		{"power of two", 16, 1},
		{"all ones", 0xFFFFFFFF, 32},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hammingWeight(tt.n)
			if got != tt.expected {
				t.Errorf("hammingWeight(%d) = %d, want %d", tt.n, got, tt.expected)
			}
		})
	}
}
