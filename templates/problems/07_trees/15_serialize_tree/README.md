# Serialize and Deserialize Binary Tree

- **Difficulty**: Hard
- **Category**: Trees
- **Topics**: binary tree, preorder traversal, string processing

## Description

Design an algorithm to serialize a binary tree into a string and deserialize that string back into the original tree structure. The serialization must capture both structure and values.

## Approach

Use preorder traversal for serialization, writing each node's value followed by a comma, and "N" for nil nodes. For deserialization, split by comma and reconstruct the tree recursively using the same preorder sequence, consuming tokens one by one.
