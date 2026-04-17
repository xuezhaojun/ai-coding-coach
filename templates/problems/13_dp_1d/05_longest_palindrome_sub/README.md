# Longest Palindromic Substring

- **Difficulty**: Medium
- **Category**: 1D Dynamic Programming
- **Topics**: dynamic programming, string, two pointers

## Description

Given a string s, return the longest palindromic substring in s.

## Approach

Use the expand-around-center technique. For each character (and each pair of adjacent characters), expand outward while characters match. Track the longest palindrome found. This achieves O(n^2) time with O(1) space, which is optimal without Manacher's algorithm.
