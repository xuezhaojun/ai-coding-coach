# Longest Common Subsequence

- **Difficulty**: Medium
- **Category**: 2D Dynamic Programming
- **Topics**: dynamic programming, string

## Description

Given two strings, return the length of their longest common subsequence. A subsequence is derived by deleting some (or no) characters without changing the order.

## Approach

Use a 1D DP array optimized from the standard 2D table. For each character in text1, iterate through text2 and track the diagonal value (prev). If characters match, dp[j] = prev + 1; otherwise, dp[j] = max(dp[j], dp[j-1]).
