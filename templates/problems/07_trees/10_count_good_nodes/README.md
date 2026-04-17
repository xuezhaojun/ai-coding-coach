# Count Good Nodes in Binary Tree

- **Difficulty**: Medium
- **Category**: Trees
- **Topics**: binary tree, DFS, recursion

## Description

Given a binary tree root, a node X is good if in the path from root to X there are no nodes with a value greater than X. Return the number of good nodes in the tree.

## Approach

Use DFS passing the maximum value seen so far along the path from root. If the current node's value is greater than or equal to the max, it is a good node. Update the max and recurse on children.
