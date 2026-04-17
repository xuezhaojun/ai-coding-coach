# Binary Tree Maximum Path Sum

- **Difficulty**: Hard
- **Category**: Trees
- **Topics**: binary tree, DFS, dynamic programming on trees

## Description

Given the root of a binary tree, return the maximum path sum of any non-empty path. A path can start and end at any node and consists of connected nodes where each node appears at most once.

## Approach

Use DFS where each call returns the maximum gain from that node going downward (single branch). At each node, compute the path sum through that node (left gain + node + right gain), updating the global max. Clamp negative gains to zero since we can choose not to extend a path.
