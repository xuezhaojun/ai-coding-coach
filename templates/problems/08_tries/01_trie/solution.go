package trie

// SolveTrie implements a prefix tree.
type SolveTrie struct {
	Root *TrieNode
}

// NewSolveTrie creates and returns a new SolveTrie.
func NewSolveTrie() SolveTrie {
	return SolveTrie{Root: &TrieNode{Children: make(map[byte]*TrieNode)}}
}

// Insert adds a word to the trie.
// Time: O(m) where m is word length, Space: O(m).
func (t *SolveTrie) Insert(word string) {
	node := t.Root
	for i := 0; i < len(word); i++ {
		c := word[i]
		if node.Children[c] == nil {
			node.Children[c] = &TrieNode{Children: make(map[byte]*TrieNode)}
		}
		node = node.Children[c]
	}
	node.IsEnd = true
}

// Search returns true if the word is in the trie.
// Time: O(m), Space: O(1).
func (t *SolveTrie) Search(word string) bool {
	node := t.findNode(word)
	return node != nil && node.IsEnd
}

// StartsWith returns true if any word in the trie starts with the given prefix.
// Time: O(m), Space: O(1).
func (t *SolveTrie) StartsWith(prefix string) bool {
	return t.findNode(prefix) != nil
}

func (t *SolveTrie) findNode(prefix string) *TrieNode {
	node := t.Root
	for i := 0; i < len(prefix); i++ {
		node = node.Children[prefix[i]]
		if node == nil {
			return nil
		}
	}
	return node
}
