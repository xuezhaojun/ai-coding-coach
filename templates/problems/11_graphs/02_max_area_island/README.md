# Max Area of Island

- **Difficulty**: Medium
- **Category**: Graphs
- **Topics**: DFS, BFS, matrix, flood fill

## Description

Given an m x n binary matrix grid, return the area of the largest island (connected component of 1s). If there is no island, return 0.

## Approach

Use DFS flood fill on each unvisited land cell. The DFS returns the count of connected cells (area). Track the maximum area found across all islands. Mark visited cells by setting them to 0.
