# Swim in Rising Water

- **Difficulty**: Hard
- **Category**: Advanced Graphs
- **Topics**: Dijkstra's algorithm, heap, binary search, BFS

## Description

Given an n x n grid where grid[i][j] represents the elevation, find the minimum time t such that you can swim from (0,0) to (n-1,n-1). At time t, you can swim to any adjacent cell with elevation <= t. You start at time grid[0][0].

## Approach

Use a modified Dijkstra's algorithm with a min-heap. The "distance" to each cell is the maximum elevation along the path. Start from (0,0) with time grid[0][0]. For each neighbor, the time is max(current time, neighbor elevation). The first time we reach (n-1,n-1) gives the optimal answer.
