# Surrounded Regions

- **Difficulty**: Medium
- **Category**: Graphs
- **Topics**: DFS, BFS, matrix, flood fill

## Description

Given an m x n matrix board containing 'X' and 'O', capture all regions that are surrounded by 'X'. A region is captured by flipping all 'O's into 'X's. Regions connected to the border cannot be captured.

## Approach

Use DFS from all border 'O' cells to mark them as temporary 'T' (safe). Then iterate through the entire board: convert remaining 'O's to 'X' (they are surrounded) and convert 'T's back to 'O' (they are border-connected and safe).
