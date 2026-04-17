package trie

// TrieNode represents a node in the Trie.
type TrieNode struct {
	Children map[byte]*TrieNode
	IsEnd    bool
}

// Trie implements a prefix tree.
type Trie struct {
	Root *TrieNode
}

// NewTrie creates and returns a new Trie.
func NewTrie() Trie {
	return Trie{}
}

// Insert adds a word to the trie.
func (t *Trie) Insert(word string) {
}

// Search returns true if the word is in the trie.
func (t *Trie) Search(word string) bool {
	return false
}

// StartsWith returns true if any word in the trie starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	return false
}
