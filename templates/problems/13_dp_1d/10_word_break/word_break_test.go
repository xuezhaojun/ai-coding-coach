package word_break

import "testing"

func TestWordBreak(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		wordDict []string
		want     bool
	}{
		{
			name:     "example leetcode",
			s:        "leetcode",
			wordDict: []string{"leet", "code"},
			want:     true,
		},
		{
			name:     "example applepenapple",
			s:        "applepenapple",
			wordDict: []string{"apple", "pen"},
			want:     true,
		},
		{
			name:     "cannot break",
			s:        "catsandog",
			wordDict: []string{"cats", "dog", "sand", "and", "cat"},
			want:     false,
		},
		{
			name:     "empty string",
			s:        "",
			wordDict: []string{"a"},
			want:     true,
		},
		{
			name:     "single char match",
			s:        "a",
			wordDict: []string{"a"},
			want:     true,
		},
		{
			name:     "single char no match",
			s:        "b",
			wordDict: []string{"a"},
			want:     false,
		},
		{
			name:     "overlapping words",
			s:        "cars",
			wordDict: []string{"car", "ca", "rs"},
			want:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := wordBreak(tt.s, tt.wordDict)
			if got != tt.want {
				t.Errorf("wordBreak(%q, %v) = %v, want %v", tt.s, tt.wordDict, got, tt.want)
			}
		})
	}
}
