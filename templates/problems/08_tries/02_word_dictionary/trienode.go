package word_dictionary

// TrieNode represents a node in the Trie.
type TrieNode struct {
	Children map[byte]*TrieNode
	IsEnd    bool
}
