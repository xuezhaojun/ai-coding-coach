# Longest Increasing Path in a Matrix

- **Difficulty**: Hard
- **Category**: 2D Dynamic Programming
- **Topics**: dynamic programming, DFS, memoization, matrix

## Description

Given an m x n matrix of integers, find the length of the longest strictly increasing path. You can move in four directions (up, down, left, right).

## Approach

Use DFS with memoization. For each cell, explore all four neighbors with strictly greater values and cache the result. Each cell is computed at most once, giving O(m*n) total time. The strictly increasing constraint prevents cycles, so no visited array is needed.
