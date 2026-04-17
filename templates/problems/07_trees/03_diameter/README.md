# Diameter of Binary Tree

- **Difficulty**: Easy
- **Category**: Trees
- **Topics**: binary tree, DFS, recursion

## Description

Given the root of a binary tree, return the length of the diameter of the tree. The diameter is the length of the longest path between any two nodes, measured by the number of edges.

## Approach

Use DFS returning the height of each subtree. At each node, the path through that node has length left_height + right_height. Track the global maximum across all nodes.
