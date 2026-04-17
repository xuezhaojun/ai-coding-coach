# Palindromic Substrings

- **Difficulty**: Medium
- **Category**: 1D Dynamic Programming
- **Topics**: dynamic programming, string, two pointers

## Description

Given a string s, return the number of palindromic substrings in it. A substring is palindromic if it reads the same forward and backward.

## Approach

Use expand-around-center for each index as an odd-length center and each adjacent pair as an even-length center. Each successful expansion adds one palindromic substring to the count.
