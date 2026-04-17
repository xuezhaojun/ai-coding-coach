# Balanced Binary Tree

- **Difficulty**: Easy
- **Category**: Trees
- **Topics**: binary tree, DFS, recursion

## Description

Given a binary tree, determine if it is height-balanced. A height-balanced binary tree is one in which the depth of the two subtrees of every node never differs by more than one.

## Approach

Use a bottom-up DFS that returns the height of each subtree, or -1 if any subtree is unbalanced. This avoids redundant height calculations and runs in O(n) time.
