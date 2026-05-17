package longest_common_sub

import "testing"

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name  string
		text1 string
		text2 string
		want  int
	}{
		{
			name:  "example 1",
			text1: "abcde",
			text2: "ace",
			want:  3,
		},
		{
			name:  "example 2",
			text1: "abc",
			text2: "abc",
			want:  3,
		},
		{
			name:  "no common",
			text1: "abc",
			text2: "def",
			want:  0,
		},
		{
			name:  "one char vs longer",
			text1: "a",
			text2: "abc",
			want:  1,
		},
		{
			name:  "single char match",
			text1: "a",
			text2: "a",
			want:  1,
		},
		{
			name:  "single char no match",
			text1: "a",
			text2: "b",
			want:  0,
		},
		{
			name:  "longer example",
			text1: "oxcpqrsvwf",
			text2: "shmtulqrypy",
			want:  2,
		},
		{
			name:  "duplicate chars must not double count",
			text1: "bsbininm",
			text2: "jmjkbkjkv",
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonSubsequence(tt.text1, tt.text2)
			if got != tt.want {
				t.Errorf("longestCommonSubsequence(%q, %q) = %d, want %d", tt.text1, tt.text2, got, tt.want)
			}
		})
	}
}
