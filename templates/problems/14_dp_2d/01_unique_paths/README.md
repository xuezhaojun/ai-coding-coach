# Unique Paths

- **Difficulty**: Medium
- **Category**: 2D Dynamic Programming
- **Topics**: dynamic programming, math, combinatorics

## Description

A robot is located at the top-left corner of an m x n grid. It can only move right or down. Count the number of unique paths to reach the bottom-right corner.

## Approach

Use a 1D DP array of size n. Each cell accumulates paths from the left (dp[j-1]) and from above (dp[j] from the previous row). This reduces space from O(m*n) to O(n) while maintaining O(m*n) time.
