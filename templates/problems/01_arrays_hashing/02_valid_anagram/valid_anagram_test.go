package valid_anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{
			name: "valid anagram",
			s:    "anagram",
			t:    "nagaram",
			want: true,
		},
		{
			name: "not anagram",
			s:    "rat",
			t:    "car",
			want: false,
		},
		{
			name: "empty strings",
			s:    "",
			t:    "",
			want: true,
		},
		{
			name: "different lengths",
			s:    "ab",
			t:    "abc",
			want: false,
		},
		{
			name: "single characters same",
			s:    "a",
			t:    "a",
			want: true,
		},
		{
			name: "single characters different",
			s:    "a",
			t:    "b",
			want: false,
		},
		{
			name: "repeated characters",
			s:    "aabb",
			t:    "bbaa",
			want: true,
		},
		{
			name: "same chars different frequency",
			s:    "aaab",
			t:    "aabb",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAnagram(tt.s, tt.t)
			if got != tt.want {
				t.Errorf("isAnagram(%q, %q) = %v, want %v", tt.s, tt.t, got, tt.want)
			}
		})
	}
}
