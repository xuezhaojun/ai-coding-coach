package edit_distance

import "testing"

func TestMinDistance(t *testing.T) {
	tests := []struct {
		name  string
		word1 string
		word2 string
		want  int
	}{
		{
			name:  "horse to ros",
			word1: "horse",
			word2: "ros",
			want:  3,
		},
		{
			name:  "intention to execution",
			word1: "intention",
			word2: "execution",
			want:  5,
		},
		{
			name:  "empty to abc",
			word1: "",
			word2: "abc",
			want:  3,
		},
		{
			name:  "abc to empty",
			word1: "abc",
			word2: "",
			want:  3,
		},
		{
			name:  "both empty",
			word1: "",
			word2: "",
			want:  0,
		},
		{
			name:  "same strings",
			word1: "abc",
			word2: "abc",
			want:  0,
		},
		{
			name:  "single char different",
			word1: "a",
			word2: "b",
			want:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minDistance(tt.word1, tt.word2)
			if got != tt.want {
				t.Errorf("minDistance(%q, %q) = %d, want %d", tt.word1, tt.word2, got, tt.want)
			}
		})
	}
}
