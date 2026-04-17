package reverse_integer

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected int
	}{
		{"positive", 123, 321},
		{"negative", -123, -321},
		{"trailing zero", 120, 21},
		{"zero", 0, 0},
		{"single digit", 5, 5},
		{"overflow positive", 1534236469, 0},
		{"overflow negative", -2147483648, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverse(tt.x)
			if got != tt.expected {
				t.Errorf("reverse(%d) = %d, want %d", tt.x, got, tt.expected)
			}
		})
	}
}
