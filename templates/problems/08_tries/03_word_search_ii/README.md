# Word Search II

- **Difficulty**: Hard
- **Category**: Tries
- **Topics**: trie, backtracking, DFS, matrix

## Description

Given an m x n board of characters and a list of words, return all words that can be constructed from letters of sequentially adjacent cells (horizontal or vertical neighbors), where each cell may only be used once per word.

## Approach

Build a trie from the word list, then DFS from each cell on the board following trie paths. Mark cells as visited during exploration to prevent reuse. When a trie node stores a complete word, add it to results and clear it to avoid duplicates.
