package permutation_in_string

import "testing"

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name     string
		s1       string
		s2       string
		expected bool
	}{
		{
			name:     "permutation exists",
			s1:       "ab",
			s2:       "eidbaooo",
			expected: true,
		},
		{
			name:     "no permutation",
			s1:       "ab",
			s2:       "eidboaoo",
			expected: false,
		},
		{
			name:     "exact match",
			s1:       "abc",
			s2:       "bca",
			expected: true,
		},
		{
			name:     "s1 longer than s2",
			s1:       "abcdef",
			s2:       "abc",
			expected: false,
		},
		{
			name:     "single character match",
			s1:       "a",
			s2:       "a",
			expected: true,
		},
		{
			name:     "single character no match",
			s1:       "a",
			s2:       "b",
			expected: false,
		},
		{
			name:     "repeated characters",
			s1:       "aab",
			s2:       "ccccbaa",
			expected: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkInclusion(tt.s1, tt.s2)
			if got != tt.expected {
				t.Errorf("checkInclusion(%q, %q) = %v, want %v", tt.s1, tt.s2, got, tt.expected)
			}
		})
	}
}
