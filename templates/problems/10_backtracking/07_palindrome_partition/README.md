# Palindrome Partitioning

- **Difficulty**: Medium
- **Category**: Backtracking
- **Topics**: backtracking, string, palindrome, dynamic programming

## Description

Given a string s, partition s such that every substring of the partition is a palindrome. Return all possible palindrome partitions.

## Approach

Use backtracking starting from index 0. At each step, try every possible end index for the current substring. If the substring is a palindrome, add it to the current partition and recurse from the next index. When the start index reaches the end of the string, record the partition.
