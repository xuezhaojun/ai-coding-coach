# Walls and Gates

- **Difficulty**: Medium
- **Category**: Graphs
- **Topics**: BFS, multi-source BFS, matrix

## Description

Given an m x n grid where -1 represents a wall, 0 represents a gate, and INF represents an empty room, fill each empty room with the distance to its nearest gate. If a room cannot reach any gate, leave it as INF.

## Approach

Use multi-source BFS starting from all gates simultaneously. Add all gate positions to the queue initially. For each cell processed, update its unvisited neighbors with distance + 1. Since BFS explores level by level, each room is assigned the shortest distance to any gate.
