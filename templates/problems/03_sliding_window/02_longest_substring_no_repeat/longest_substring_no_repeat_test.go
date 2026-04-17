package longest_substring_no_repeat

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "basic case",
			s:        "abcabcbb",
			expected: 3,
		},
		{
			name:     "all same characters",
			s:        "bbbbb",
			expected: 1,
		},
		{
			name:     "mixed repeats",
			s:        "pwwkew",
			expected: 3,
		},
		{
			name:     "empty string",
			s:        "",
			expected: 0,
		},
		{
			name:     "single character",
			s:        "a",
			expected: 1,
		},
		{
			name:     "all unique",
			s:        "abcdef",
			expected: 6,
		},
		{
			name:     "spaces and special chars",
			s:        "a b c",
			expected: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.s)
			if got != tt.expected {
				t.Errorf("lengthOfLongestSubstring(%q) = %d, want %d", tt.s, got, tt.expected)
			}
		})
	}
}
