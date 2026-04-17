# Partition Labels

- **Difficulty**: Medium
- **Category**: Greedy
- **Topics**: string, greedy, hash table, two pointers

## Description

Given a string s, partition it into as many parts as possible so that each letter appears in at most one part. Return a list of integers representing the size of each part.

## Approach

First pass: record the last occurrence index of each character. Second pass: greedily extend the current partition's end to include the last occurrence of each character encountered. When the current index equals the partition end, cut the partition.
