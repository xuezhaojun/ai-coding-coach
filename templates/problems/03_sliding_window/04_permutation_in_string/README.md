# Permutation in String

- **Difficulty**: Medium
- **Category**: Sliding Window
- **Topics**: string, hash table, sliding window, two pointers

## Description

Given two strings s1 and s2, return true if s2 contains a permutation of s1. In other words, return true if one of s1's permutations is a substring of s2.

## Approach

Use a fixed-size sliding window of length len(s1) over s2. Maintain character frequency arrays for both s1 and the current window. Slide the window one character at a time, adding the new character and removing the old one, then compare the two arrays. Array comparison in Go is O(1) for fixed-size arrays.
