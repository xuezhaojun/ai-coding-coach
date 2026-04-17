package word_dictionary

// SolveWordDictionary supports adding words and searching with '.' wildcards.
type SolveWordDictionary struct {
	Root *TrieNode
}

// NewSolveWordDictionary creates and returns a new SolveWordDictionary.
func NewSolveWordDictionary() SolveWordDictionary {
	return SolveWordDictionary{Root: &TrieNode{Children: make(map[byte]*TrieNode)}}
}

// AddWord adds a word to the dictionary.
// Time: O(m), Space: O(m).
func (wd *SolveWordDictionary) AddWord(word string) {
	node := wd.Root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if node.Children[c] == nil {
			node.Children[c] = &TrieNode{Children: make(map[byte]*TrieNode)}
		}
		node = node.Children[c]
	}
	node.IsEnd = true
}

// Search returns true if the word matches any string in the dictionary.
// '.' can match any single character.
// Time: O(26^d * m) worst case where d is number of dots, Space: O(m).
func (wd *SolveWordDictionary) Search(word string) bool {
	return searchNode(wd.Root, word, 0)
}

func searchNode(node *TrieNode, word string, i int) bool {
	if i == len(word) {
		return node.IsEnd
	}
	c := word[i]
	if c == '.' {
		for _, child := range node.Children {
			if searchNode(child, word, i+1) {
				return true
			}
		}
		return false
	}
	child := node.Children[c]
	if child == nil {
		return false
	}
	return searchNode(child, word, i+1)
}
