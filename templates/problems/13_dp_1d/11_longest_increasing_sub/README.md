# Longest Increasing Subsequence

- **Difficulty**: Medium
- **Category**: 1D Dynamic Programming
- **Topics**: dynamic programming, binary search

## Description

Given an integer array, return the length of the longest strictly increasing subsequence.

## Approach

Maintain a "tails" array where tails[i] holds the smallest tail element of all increasing subsequences of length i+1. For each number, use binary search to find its position. If it extends the longest subsequence, append it; otherwise, replace the element at the found position. The length of the tails array is the answer.
