# Binary Tree Level Order Traversal

- **Difficulty**: Medium
- **Category**: Trees
- **Topics**: binary tree, BFS, queue

## Description

Given the root of a binary tree, return the level order traversal of its nodes' values (i.e., from left to right, level by level).

## Approach

Use BFS with a queue. Process nodes level by level: for each level, dequeue all current nodes, record their values, and enqueue their children. Each level forms one subarray in the result.
