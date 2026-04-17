# Maximum Depth of Binary Tree

- **Difficulty**: Easy
- **Category**: Trees
- **Topics**: binary tree, recursion, DFS

## Description

Given the root of a binary tree, return its maximum depth. The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

## Approach

Use DFS recursion: the depth of a tree is 1 plus the maximum of the depths of its left and right subtrees. Base case: a nil node has depth 0.
