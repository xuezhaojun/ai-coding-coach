# Word Ladder

- **Difficulty**: Hard
- **Category**: Graphs
- **Topics**: BFS, string, hash set

## Description

Given two words beginWord and endWord and a dictionary wordList, return the number of words in the shortest transformation sequence from beginWord to endWord, where each transformed word must exist in wordList. Each step changes exactly one letter. Return 0 if no such sequence exists.

## Approach

Use BFS starting from beginWord. At each level, try changing every character position to every letter a-z. If the new word exists in the word set and has not been visited, add it to the queue. BFS guarantees the first time endWord is reached, it is via the shortest path. Return the level count + 1.
