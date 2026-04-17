# Binary Tree Right Side View

- **Difficulty**: Medium
- **Category**: Trees
- **Topics**: binary tree, BFS, queue

## Description

Given the root of a binary tree, return the values of the nodes you can see when looking at the tree from the right side, ordered from top to bottom.

## Approach

Use BFS level-by-level traversal. For each level, the last node processed is the rightmost node visible from the right side. Append that node's value to the result.
