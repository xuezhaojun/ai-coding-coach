package reverse_bits

import "testing"

func TestReverseBits(t *testing.T) {
	tests := []struct {
		name     string
		num      uint32
		expected uint32
	}{
		{"example 1", 0b00000010100101000001111010011100, 964176192},
		{"example 2", 0b11111111111111111111111111111101, 3221225471},
		{"zero", 0, 0},
		{"all ones", 0xFFFFFFFF, 0xFFFFFFFF},
		{"one", 1, 0x80000000},
		{"power of two", 0x80000000, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBits(tt.num)
			if got != tt.expected {
				t.Errorf("reverseBits(%d) = %d, want %d", tt.num, got, tt.expected)
			}
		})
	}
}
