# Same Tree

- **Difficulty**: Easy
- **Category**: Trees
- **Topics**: binary tree, DFS, recursion

## Description

Given the roots of two binary trees, check if they are the same. Two binary trees are the same if they are structurally identical and all corresponding nodes have the same value.

## Approach

Recursively compare both trees simultaneously. If both nodes are nil, they match. If one is nil or values differ, they don't match. Otherwise recurse on left and right children.
