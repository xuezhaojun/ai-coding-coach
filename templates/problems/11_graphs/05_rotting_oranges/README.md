# Rotting Oranges

- **Difficulty**: Medium
- **Category**: Graphs
- **Topics**: BFS, multi-source BFS, matrix

## Description

Given an m x n grid where 0 is empty, 1 is a fresh orange, and 2 is a rotten orange, return the minimum number of minutes that must elapse until no cell has a fresh orange. If impossible, return -1. Each minute, rotten oranges rot adjacent fresh oranges.

## Approach

Use multi-source BFS starting from all initially rotten oranges. Process level by level where each level represents one minute. Track the count of fresh oranges and decrement as they rot. After BFS completes, return -1 if any fresh oranges remain, otherwise return the number of minutes elapsed.
