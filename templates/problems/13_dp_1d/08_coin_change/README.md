# Coin Change

- **Difficulty**: Medium
- **Category**: 1D Dynamic Programming
- **Topics**: dynamic programming, breadth-first search

## Description

Given an array of coin denominations and a target amount, return the fewest number of coins needed to make that amount. Return -1 if it cannot be made.

## Approach

Use bottom-up DP where dp[i] represents the minimum coins needed to make amount i. Initialize all values to amount+1 (impossible sentinel). For each amount, try every coin and take the minimum. If dp[amount] exceeds amount, return -1.
