package word_search

import "testing"

func TestExist(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{
			name:  "word exists",
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCCED",
			want:  true,
		},
		{
			name:  "word exists path SEE",
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "SEE",
			want:  true,
		},
		{
			name:  "word does not exist",
			board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word:  "ABCB",
			want:  false,
		},
		{
			name:  "single cell match",
			board: [][]byte{{'A'}},
			word:  "A",
			want:  true,
		},
		{
			name:  "single cell no match",
			board: [][]byte{{'A'}},
			word:  "B",
			want:  false,
		},
		{
			name:  "word longer than board cells",
			board: [][]byte{{'A', 'B'}, {'C', 'D'}},
			word:  "ABCDA",
			want:  false,
		},
		{
			name:  "snake path",
			board: [][]byte{{'A', 'B', 'C'}, {'D', 'E', 'F'}, {'G', 'H', 'I'}},
			word:  "ABCFEDGHI",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := exist(tt.board, tt.word)
			if got != tt.want {
				t.Errorf("exist(%v, %q) = %v, want %v", tt.board, tt.word, got, tt.want)
			}
		})
	}
}
