# Interleaving String

- **Difficulty**: Medium
- **Category**: 2D Dynamic Programming
- **Topics**: dynamic programming, string

## Description

Given strings s1, s2, and s3, determine whether s3 is formed by interleaving s1 and s2 while preserving the relative order of characters in each.

## Approach

Use a 1D DP array where dp[j] indicates if s3[0:i+j] can be formed by interleaving s1[0:i] and s2[0:j]. At each cell, check if the current character of s3 matches either s1[i-1] (from above) or s2[j-1] (from left).
