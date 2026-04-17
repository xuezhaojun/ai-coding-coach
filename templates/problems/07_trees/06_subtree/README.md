# Subtree of Another Tree

- **Difficulty**: Easy
- **Category**: Trees
- **Topics**: binary tree, DFS, recursion

## Description

Given the roots of two binary trees root and subRoot, return true if there is a subtree of root with the same structure and node values as subRoot.

## Approach

For each node in root, check if the subtree rooted there is identical to subRoot using a same-tree comparison. Recurse through the main tree trying each node as a potential match point.
