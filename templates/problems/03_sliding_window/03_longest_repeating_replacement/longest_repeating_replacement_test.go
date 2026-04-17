package longest_repeating_replacement

import "testing"

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "basic case",
			s:        "ABAB",
			k:        2,
			expected: 4,
		},
		{
			name:     "replace one",
			s:        "AABABBA",
			k:        1,
			expected: 4,
		},
		{
			name:     "no replacement needed",
			s:        "AAAA",
			k:        0,
			expected: 4,
		},
		{
			name:     "single character",
			s:        "A",
			k:        0,
			expected: 1,
		},
		{
			name:     "k equals string length",
			s:        "ABCD",
			k:        4,
			expected: 4,
		},
		{
			name:     "alternating with k=0",
			s:        "ABABAB",
			k:        0,
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := characterReplacement(tt.s, tt.k)
			if got != tt.expected {
				t.Errorf("characterReplacement(%q, %d) = %d, want %d", tt.s, tt.k, got, tt.expected)
			}
		})
	}
}
