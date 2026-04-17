# Invert Binary Tree

- **Difficulty**: Easy
- **Category**: Trees
- **Topics**: binary tree, recursion, DFS

## Description

Given the root of a binary tree, invert the tree (mirror it) and return its root. Every left child is swapped with its right child at each node.

## Approach

Recursively swap the left and right children of each node. Base case is a nil node. After swapping at the current node, recurse on both subtrees.
