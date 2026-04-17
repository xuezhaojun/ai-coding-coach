package word_dictionary

import "testing"

func TestWordDictionary(t *testing.T) {
	tests := []struct {
		name     string
		adds     []string
		searches []struct{ word string; want bool }
	}{
		{
			name: "basic with wildcards",
			adds: []string{"bad", "dad", "mad"},
			searches: []struct{ word string; want bool }{
				{"pad", false},
				{"bad", true},
				{".ad", true},
				{"b..", true},
				{"...", true},
				{"....", false},
			},
		},
		{
			name: "exact match only",
			adds: []string{"hello"},
			searches: []struct{ word string; want bool }{
				{"hello", true},
				{"hell", false},
				{"helloo", false},
			},
		},
		{
			name: "all wildcards",
			adds: []string{"ab", "cd"},
			searches: []struct{ word string; want bool }{
				{"..", true},
				{".", false},
				{"...", false},
			},
		},
		{
			name: "single char",
			adds: []string{"a"},
			searches: []struct{ word string; want bool }{
				{".", true},
				{"a", true},
				{"..", false},
			},
		},
		{
			name: "mixed wildcards",
			adds: []string{"apple", "ample"},
			searches: []struct{ word string; want bool }{
				{"a.ple", true},
				{"a..le", true},
				{"a...e", true},
				{"a....", true},
				{"b....", false},
			},
		},
		{
			name: "no words added",
			adds: []string{},
			searches: []struct{ word string; want bool }{
				{"a", false},
				{".", false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wd := NewWordDictionary()
			for _, w := range tt.adds {
				wd.AddWord(w)
			}
			for _, s := range tt.searches {
				if got := wd.Search(s.word); got != s.want {
					t.Errorf("Search(%q) = %v, want %v", s.word, got, s.want)
				}
			}
		})
	}
}
