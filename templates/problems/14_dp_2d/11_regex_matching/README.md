# Regular Expression Matching

- **Difficulty**: Hard
- **Category**: 2D Dynamic Programming
- **Topics**: dynamic programming, string, recursion

## Description

Implement regular expression matching with support for '.' (matches any single character) and '*' (matches zero or more of the preceding element). The matching should cover the entire input string.

## Approach

Use space-optimized DP with a 1D array. For '*', either match zero occurrences (dp[j-2]) or one-or-more if the preceding character matches the current string character (use the value from the previous row). For '.' or exact match, use the diagonal value. Track the diagonal with a prev variable.
