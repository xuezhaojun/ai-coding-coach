package valid_palindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "alphanumeric palindrome with spaces and punctuation",
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "not a palindrome",
			s:        "race a car",
			expected: false,
		},
		{
			name:     "empty string is palindrome",
			s:        " ",
			expected: true,
		},
		{
			name:     "single character",
			s:        "a",
			expected: true,
		},
		{
			name:     "mixed case palindrome",
			s:        "Aa",
			expected: true,
		},
		{
			name:     "digits in palindrome",
			s:        "0P",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPalindrome(tt.s)
			if got != tt.expected {
				t.Errorf("isPalindrome(%q) = %v, want %v", tt.s, got, tt.expected)
			}
		})
	}
}
