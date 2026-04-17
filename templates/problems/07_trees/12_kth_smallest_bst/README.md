# Kth Smallest Element in a BST

- **Difficulty**: Medium
- **Category**: Trees
- **Topics**: binary search tree, inorder traversal, DFS

## Description

Given the root of a binary search tree and an integer k, return the kth smallest value (1-indexed) of all the values in the BST.

## Approach

Perform an inorder traversal of the BST, which visits nodes in ascending order. Count visited nodes and stop as soon as the kth node is reached. This runs in O(h + k) time.
