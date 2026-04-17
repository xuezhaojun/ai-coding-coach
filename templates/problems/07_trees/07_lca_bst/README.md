# Lowest Common Ancestor of BST

- **Difficulty**: Medium
- **Category**: Trees
- **Topics**: binary search tree, BST properties, iteration

## Description

Given a binary search tree and two nodes p and q, find their lowest common ancestor (LCA). The LCA is the deepest node that is an ancestor of both p and q.

## Approach

Leverage BST properties: if both p and q are smaller than current node, go left; if both are larger, go right. Otherwise the current node is the split point and thus the LCA. Runs in O(h) time with O(1) space using iteration.
