# Jump Game II

- **Difficulty**: Medium
- **Category**: Greedy
- **Topics**: array, greedy, BFS

## Description

Given an integer array nums where each element represents the maximum jump length from that position, return the minimum number of jumps to reach the last index. You can assume you can always reach the last index.

## Approach

Use a greedy BFS-like approach: maintain the current jump's boundary and the farthest reachable position. When you reach the current boundary, increment jumps and extend the boundary to the farthest point seen.
