# Number of Islands

- **Difficulty**: Medium
- **Category**: Graphs
- **Topics**: DFS, BFS, matrix, flood fill

## Description

Given an m x n 2D binary grid which represents a map of '1's (land) and '0's (water), return the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically.

## Approach

Iterate through each cell. When a '1' is found, increment the island count and use DFS to mark all connected land cells as '0' (visited). This flood-fill ensures each island is counted exactly once.
