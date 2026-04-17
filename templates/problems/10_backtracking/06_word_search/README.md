# Word Search

- **Difficulty**: Medium
- **Category**: Backtracking
- **Topics**: backtracking, DFS, matrix

## Description

Given an m x n grid of characters board and a string word, return true if word exists in the grid. The word can be constructed from letters of sequentially adjacent cells (horizontally or vertically), and each cell may be used at most once.

## Approach

For each cell matching the first character, start a DFS. Mark visited cells by temporarily replacing them with a sentinel character. Explore all four directions for the next character. Restore the cell on backtrack. Return true as soon as any path matches the full word.
