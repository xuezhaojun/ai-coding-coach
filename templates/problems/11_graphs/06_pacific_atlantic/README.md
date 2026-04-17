# Pacific Atlantic Water Flow

- **Difficulty**: Medium
- **Category**: Graphs
- **Topics**: DFS, BFS, matrix

## Description

Given an m x n matrix of heights, find all cells where water can flow to both the Pacific (top/left edges) and Atlantic (bottom/right edges) oceans. Water flows from a cell to adjacent cells with equal or lower height.

## Approach

Use reverse DFS from ocean borders. Start DFS from all Pacific-border cells and all Atlantic-border cells separately, flowing uphill (to cells with >= height). Maintain two boolean matrices. Cells reachable from both oceans are in the result.
