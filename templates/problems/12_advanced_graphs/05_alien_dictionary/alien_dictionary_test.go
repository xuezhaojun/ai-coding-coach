package alien_dictionary

import "testing"

func TestAlienOrder(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  string
	}{
		{
			name:  "standard order",
			words: []string{"wrt", "wrf", "er", "ett", "rftt"},
			want:  "wertf",
		},
		{
			name:  "simple two words",
			words: []string{"z", "x"},
			want:  "zx",
		},
		{
			name:  "invalid order",
			words: []string{"z", "x", "z"},
			want:  "",
		},
		{
			name:  "single word",
			words: []string{"abc"},
			want:  "abc",
		},
		{
			name:  "prefix violation",
			words: []string{"abc", "ab"},
			want:  "",
		},
		{
			name:  "single characters",
			words: []string{"z", "z"},
			want:  "z",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := alienOrder(tt.words)
			if tt.want == "" {
				if got != "" {
					t.Errorf("alienOrder(%v) = %q, want empty string", tt.words, got)
				}
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("alienOrder(%v) = %q (len %d), want len %d", tt.words, got, len(got), len(tt.want))
				return
			}
			// verify the ordering is consistent with input
			charPos := make(map[byte]int)
			for i := 0; i < len(got); i++ {
				charPos[got[i]] = i
			}
			for i := 0; i < len(tt.words)-1; i++ {
				w1, w2 := tt.words[i], tt.words[i+1]
				for j := 0; j < len(w1) && j < len(w2); j++ {
					if w1[j] != w2[j] {
						if charPos[w1[j]] > charPos[w2[j]] {
							t.Errorf("alienOrder(%v) = %q violates ordering of %q before %q", tt.words, got, tt.words[i], tt.words[i+1])
						}
						break
					}
				}
			}
		})
	}
}
