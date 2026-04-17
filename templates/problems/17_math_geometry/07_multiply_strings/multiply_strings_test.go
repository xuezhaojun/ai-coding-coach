package multiply_strings

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct {
		name     string
		num1     string
		num2     string
		expected string
	}{
		{"example 1", "2", "3", "6"},
		{"example 2", "123", "456", "56088"},
		{"multiply by zero", "0", "52", "0"},
		{"both zeros", "0", "0", "0"},
		{"single digits", "9", "9", "81"},
		{"large numbers", "999", "999", "998001"},
		{"one and number", "1", "12345", "12345"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := multiply(tt.num1, tt.num2)
			if got != tt.expected {
				t.Errorf("multiply(%q, %q) = %q, want %q", tt.num1, tt.num2, got, tt.expected)
			}
		})
	}
}
