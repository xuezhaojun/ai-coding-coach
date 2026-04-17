# Construct Binary Tree from Preorder and Inorder

- **Difficulty**: Medium
- **Category**: Trees
- **Topics**: binary tree, recursion, hash map, divide and conquer

## Description

Given two integer arrays preorder and inorder where preorder is the preorder traversal and inorder is the inorder traversal of the same tree, construct and return the binary tree.

## Approach

The first element of preorder is always the root. Find its index in inorder to split into left and right subtrees. Use a hash map for O(1) index lookup, and a global pointer into preorder to pick roots in sequence. Recurse on left then right subtrees.
