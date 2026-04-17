package longest_palindrome_sub

import "testing"

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string // multiple valid answers
	}{
		{
			name: "example babad",
			s:    "babad",
			want: []string{"bab", "aba"},
		},
		{
			name: "example cbbd",
			s:    "cbbd",
			want: []string{"bb"},
		},
		{
			name: "single character",
			s:    "a",
			want: []string{"a"},
		},
		{
			name: "all same characters",
			s:    "aaaa",
			want: []string{"aaaa"},
		},
		{
			name: "entire string is palindrome",
			s:    "racecar",
			want: []string{"racecar"},
		},
		{
			name: "no repeats",
			s:    "abcde",
			want: []string{"a", "b", "c", "d", "e"},
		},
		{
			name: "two characters same",
			s:    "bb",
			want: []string{"bb"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestPalindrome(tt.s)
			valid := false
			for _, w := range tt.want {
				if got == w {
					valid = true
					break
				}
			}
			if !valid {
				t.Errorf("longestPalindrome(%q) = %q, want one of %v", tt.s, got, tt.want)
			}
		})
	}
}
