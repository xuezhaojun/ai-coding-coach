package word_ladder

import "testing"

func TestLadderLength(t *testing.T) {
	tests := []struct {
		name      string
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		{
			name:      "standard transformation",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			want:      5,
		},
		{
			name:      "no valid transformation",
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			want:      0,
		},
		{
			name:      "one step",
			beginWord: "hot",
			endWord:   "dot",
			wordList:  []string{"dot"},
			want:      2,
		},
		{
			name:      "same begin and end",
			beginWord: "hit",
			endWord:   "hit",
			wordList:  []string{"hit"},
			want:      0,
		},
		{
			name:      "end word not in list",
			beginWord: "abc",
			endWord:   "xyz",
			wordList:  []string{"abc", "xbc"},
			want:      0,
		},
		{
			name:      "longer path",
			beginWord: "a",
			endWord:   "c",
			wordList:  []string{"a", "b", "c"},
			want:      2,
		},
		{
			name:      "two letter words",
			beginWord: "ab",
			endWord:   "cd",
			wordList:  []string{"ad", "cb", "cd"},
			want:      3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ladderLength(tt.beginWord, tt.endWord, tt.wordList)
			if got != tt.want {
				t.Errorf("ladderLength(%q, %q, %v) = %v, want %v",
					tt.beginWord, tt.endWord, tt.wordList, got, tt.want)
			}
		})
	}
}
