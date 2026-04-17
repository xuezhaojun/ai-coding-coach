package happy_number

import "testing"

func TestIsHappy(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected bool
	}{
		{"happy 19", 19, true},
		{"happy 1", 1, true},
		{"not happy 2", 2, false},
		{"happy 7", 7, true},
		{"not happy 4", 4, false},
		{"happy 100", 100, true},
		{"not happy 20", 20, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isHappy(tt.n)
			if got != tt.expected {
				t.Errorf("isHappy(%d) = %v, want %v", tt.n, got, tt.expected)
			}
		})
	}
}
