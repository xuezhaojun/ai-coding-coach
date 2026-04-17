# Word Break

- **Difficulty**: Medium
- **Category**: 1D Dynamic Programming
- **Topics**: dynamic programming, string, hash table

## Description

Given a string s and a dictionary of words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

## Approach

Use a boolean DP array where dp[i] indicates whether s[0:i] can be segmented. For each position i, check all possible split points j; if dp[j] is true and s[j:i] is in the dictionary, set dp[i] to true. Use a hash set for O(1) dictionary lookups.
