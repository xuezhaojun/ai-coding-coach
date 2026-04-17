# Decode Ways

- **Difficulty**: Medium
- **Category**: 1D Dynamic Programming
- **Topics**: dynamic programming, string

## Description

A message containing letters A-Z is encoded as numbers 1-26. Given a string of digits, return the number of ways to decode it.

## Approach

Use bottom-up DP with two variables. At each position, check if the current digit is valid (non-zero) and if the two-digit number formed with the previous digit is between 10 and 26. Sum the valid decodings accordingly.
