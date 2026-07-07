package distinct_subsequences

import "testing"

func TestNumDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		tt   string
		want int
	}{
		{
			name: "example rabbbit",
			s:    "rabbbit",
			tt:   "rabbit",
			want: 3,
		},
		{
			name: "example babgbag",
			s:    "babgbag",
			tt:   "bag",
			want: 5,
		},
		{
			name: "no match",
			s:    "abc",
			tt:   "def",
			want: 0,
		},
		{
			name: "t longer than s",
			s:    "ab",
			tt:   "abc",
			want: 0,
		},
		{
			name: "equal strings",
			s:    "abc",
			tt:   "abc",
			want: 1,
		},
		{
			name: "empty t",
			s:    "abc",
			tt:   "",
			want: 1,
		},
		{
			name: "single char repeated",
			s:    "aaa",
			tt:   "a",
			want: 3,
		},
	}

	for _, tt2 := range tests {
		t.Run(tt2.name, func(t *testing.T) {
			got := numDistinct(tt2.s, tt2.tt)
			if got != tt2.want {
				t.Errorf("numDistinct(%q, %q) = %d, want %d", tt2.s, tt2.tt, got, tt2.want)
			}
		})
	}
}
