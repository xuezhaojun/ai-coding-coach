package palindromic_substrings

import "testing"

func TestCountSubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example abc",
			s:    "abc",
			want: 3,
		},
		{
			name: "example aaa",
			s:    "aaa",
			want: 6,
		},
		{
			name: "single char",
			s:    "a",
			want: 1,
		},
		{
			name: "two same chars",
			s:    "aa",
			want: 3,
		},
		{
			name: "two different chars",
			s:    "ab",
			want: 2,
		},
		{
			name: "racecar",
			s:    "racecar",
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countSubstrings(tt.s)
			if got != tt.want {
				t.Errorf("countSubstrings(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
