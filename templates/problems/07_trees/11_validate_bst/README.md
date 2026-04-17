# Validate Binary Search Tree

- **Difficulty**: Medium
- **Category**: Trees
- **Topics**: binary search tree, DFS, recursion

## Description

Given the root of a binary tree, determine if it is a valid binary search tree (BST). A valid BST requires that every node's value is strictly between the allowed minimum and maximum for its position.

## Approach

Use DFS passing valid range (min, max) down the tree. For the left child, update max to the current node's value; for the right child, update min. If any node violates its range, return false.
