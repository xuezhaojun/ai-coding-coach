# Distinct Subsequences

- **Difficulty**: Hard
- **Category**: 2D Dynamic Programming
- **Topics**: dynamic programming, string

## Description

Given two strings s and t, return the number of distinct subsequences of s that equal t.

## Approach

Use a 1D DP array where dp[j] represents the number of ways to form t[0:j] from the characters of s seen so far. Iterate backwards through j to avoid using updated values. When s[i-1] == t[j-1], add dp[j-1] to dp[j].
