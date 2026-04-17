package regex_matching

import "testing"

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name string
		s    string
		p    string
		want bool
	}{
		{
			name: "no match",
			s:    "aa",
			p:    "a",
			want: false,
		},
		{
			name: "star matches multiple",
			s:    "aa",
			p:    "a*",
			want: true,
		},
		{
			name: "dot star matches all",
			s:    "ab",
			p:    ".*",
			want: true,
		},
		{
			name: "mixed pattern",
			s:    "aab",
			p:    "c*a*b",
			want: true,
		},
		{
			name: "empty string empty pattern",
			s:    "",
			p:    "",
			want: true,
		},
		{
			name: "empty string star pattern",
			s:    "",
			p:    "a*",
			want: true,
		},
		{
			name: "complex pattern",
			s:    "mississippi",
			p:    "mis*is*ip*.",
			want: true,
		},
		{
			name: "dot matches single",
			s:    "ab",
			p:    ".b",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isMatch(tt.s, tt.p)
			if got != tt.want {
				t.Errorf("isMatch(%q, %q) = %v, want %v", tt.s, tt.p, got, tt.want)
			}
		})
	}
}
