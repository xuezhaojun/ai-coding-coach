package interleaving_string

import "testing"

func TestIsInterleave(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		s3   string
		want bool
	}{
		{
			name: "example true",
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbcbcac",
			want: true,
		},
		{
			name: "example false",
			s1:   "aabcc",
			s2:   "dbbca",
			s3:   "aadbbbaccc",
			want: false,
		},
		{
			name: "both empty",
			s1:   "",
			s2:   "",
			s3:   "",
			want: true,
		},
		{
			name: "s1 empty",
			s1:   "",
			s2:   "abc",
			s3:   "abc",
			want: true,
		},
		{
			name: "s2 empty",
			s1:   "abc",
			s2:   "",
			s3:   "abc",
			want: true,
		},
		{
			name: "length mismatch",
			s1:   "a",
			s2:   "b",
			s3:   "abc",
			want: false,
		},
		{
			name: "single chars true",
			s1:   "a",
			s2:   "b",
			s3:   "ab",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isInterleave(tt.s1, tt.s2, tt.s3)
			if got != tt.want {
				t.Errorf("isInterleave(%q, %q, %q) = %v, want %v", tt.s1, tt.s2, tt.s3, got, tt.want)
			}
		})
	}
}
