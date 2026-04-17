# Design Add and Search Words Data Structure

- **Difficulty**: Medium
- **Category**: Tries
- **Topics**: trie, DFS, backtracking, wildcard matching

## Description

Design a data structure that supports adding new words and searching for words where '.' can match any single letter. Implement addWord and search operations.

## Approach

Use a trie for storage. AddWord is standard trie insertion. For search, traverse the trie character by character. When encountering '.', branch out and try all children recursively. Return true if any branch leads to a complete match.
