package decode_ways

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "example 12",
			s:    "12",
			want: 2,
		},
		{
			name: "example 226",
			s:    "226",
			want: 3,
		},
		{
			name: "leading zero",
			s:    "06",
			want: 0,
		},
		{
			name: "single digit",
			s:    "1",
			want: 1,
		},
		{
			name: "example 11106",
			s:    "11106",
			want: 2,
		},
		{
			name: "all ones",
			s:    "1111",
			want: 5,
		},
		{
			name: "example 10",
			s:    "10",
			want: 1,
		},
		{
			name: "example 27",
			s:    "27",
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numDecodings(tt.s)
			if got != tt.want {
				t.Errorf("numDecodings(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
