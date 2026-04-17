# Alien Dictionary

- **Difficulty**: Hard
- **Category**: Advanced Graphs
- **Topics**: topological sort, BFS, graph, string

## Description

Given a sorted list of words in an alien language, derive the order of characters in the alien alphabet. If the order is invalid or cannot be determined, return an empty string.

## Approach

Compare adjacent words to extract ordering constraints (the first differing character gives an edge). Build a directed graph and use BFS-based topological sort (Kahn's algorithm) to produce the character order. Detect invalid cases: prefix violations (longer word before its prefix) and cycles (result length does not match total unique characters).
