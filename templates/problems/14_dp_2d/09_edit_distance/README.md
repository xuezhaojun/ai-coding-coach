# Edit Distance

- **Difficulty**: Medium
- **Category**: 2D Dynamic Programming
- **Topics**: dynamic programming, string

## Description

Given two strings, return the minimum number of operations (insert, delete, replace) to convert one string into the other.

## Approach

Use space-optimized DP with a 1D array. Track the diagonal value (prev) to simulate the 2D table. If characters match, no operation is needed (dp[j] = prev). Otherwise, take 1 + minimum of insert (dp[j-1]), delete (dp[j]), or replace (prev).
