package pow_x_n

import (
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	tests := []struct {
		name     string
		x        float64
		n        int
		expected float64
	}{
		{"positive exponent", 2.0, 10, 1024.0},
		{"negative exponent", 2.0, -2, 0.25},
		{"zero exponent", 5.0, 0, 1.0},
		{"exponent 1", 3.0, 1, 3.0},
		{"fractional base", 2.1, 3, 9.261},
		{"base 1", 1.0, 100, 1.0},
		{"base 0", 0.0, 5, 0.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := myPow(tt.x, tt.n)
			if math.Abs(got-tt.expected) > 1e-3 {
				t.Errorf("myPow(%v, %d) = %v, want %v", tt.x, tt.n, got, tt.expected)
			}
		})
	}
}
