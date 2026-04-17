# Implement Trie / Prefix Tree

- **Difficulty**: Medium
- **Category**: Tries
- **Topics**: trie, prefix tree, hash map, string

## Description

Implement a trie (prefix tree) with insert, search, and startsWith methods. Insert adds a word, search checks if a word exists, and startsWith checks if any word has a given prefix.

## Approach

Use a tree of TrieNode objects where each node has a map of children keyed by character and a boolean marking word endings. Traverse the trie character by character for all operations.
