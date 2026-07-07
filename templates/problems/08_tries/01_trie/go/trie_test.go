package trie

import "testing"

func TestTrie(t *testing.T) {
	tests := []struct {
		name       string
		inserts    []string
		searches   []struct{ word string; want bool }
		prefixes   []struct{ prefix string; want bool }
	}{
		{
			name:    "basic operations",
			inserts: []string{"apple"},
			searches: []struct{ word string; want bool }{
				{"apple", true},
				{"app", false},
				{"apples", false},
			},
			prefixes: []struct{ prefix string; want bool }{
				{"app", true},
				{"apple", true},
				{"b", false},
			},
		},
		{
			name:    "multiple words",
			inserts: []string{"apple", "app", "banana"},
			searches: []struct{ word string; want bool }{
				{"apple", true},
				{"app", true},
				{"banana", true},
				{"ban", false},
			},
			prefixes: []struct{ prefix string; want bool }{
				{"app", true},
				{"ban", true},
				{"cat", false},
			},
		},
		{
			name:    "empty string",
			inserts: []string{""},
			searches: []struct{ word string; want bool }{
				{"", true},
				{"a", false},
			},
			prefixes: []struct{ prefix string; want bool }{
				{"", true},
			},
		},
		{
			name:    "single char words",
			inserts: []string{"a", "b", "c"},
			searches: []struct{ word string; want bool }{
				{"a", true},
				{"b", true},
				{"d", false},
			},
			prefixes: []struct{ prefix string; want bool }{
				{"a", true},
				{"d", false},
			},
		},
		{
			name:    "overlapping words",
			inserts: []string{"the", "then", "them", "there"},
			searches: []struct{ word string; want bool }{
				{"the", true},
				{"then", true},
				{"they", false},
				{"there", true},
			},
			prefixes: []struct{ prefix string; want bool }{
				{"the", true},
				{"th", true},
				{"thx", false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			trie := NewTrie()
			for _, w := range tt.inserts {
				trie.Insert(w)
			}
			for _, s := range tt.searches {
				if got := trie.Search(s.word); got != s.want {
					t.Errorf("Search(%q) = %v, want %v", s.word, got, s.want)
				}
			}
			for _, p := range tt.prefixes {
				if got := trie.StartsWith(p.prefix); got != p.want {
					t.Errorf("StartsWith(%q) = %v, want %v", p.prefix, got, p.want)
				}
			}
		})
	}
}
