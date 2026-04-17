package min_window_substring

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t_str    string
		expected string
	}{
		{
			name:     "basic case",
			s:        "ADOBECODEBANC",
			t_str:    "ABC",
			expected: "BANC",
		},
		{
			name:     "exact match",
			s:        "a",
			t_str:    "a",
			expected: "a",
		},
		{
			name:     "no valid window",
			s:        "a",
			t_str:    "aa",
			expected: "",
		},
		{
			name:     "t not in s",
			s:        "abc",
			t_str:    "z",
			expected: "",
		},
		{
			name:     "entire string is window",
			s:        "abc",
			t_str:    "abc",
			expected: "abc",
		},
		{
			name:     "duplicate chars in t",
			s:        "aaabbc",
			t_str:    "aab",
			expected: "aab",
		},
		{
			name:     "empty s",
			s:        "",
			t_str:    "a",
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minWindow(tt.s, tt.t_str)
			if got != tt.expected {
				t.Errorf("minWindow(%q, %q) = %q, want %q", tt.s, tt.t_str, got, tt.expected)
			}
		})
	}
}
