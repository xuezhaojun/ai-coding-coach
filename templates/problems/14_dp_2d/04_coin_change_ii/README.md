# Coin Change II

- **Difficulty**: Medium
- **Category**: 2D Dynamic Programming
- **Topics**: dynamic programming, unbounded knapsack

## Description

Given an amount and an array of coin denominations, return the number of combinations that make up the amount. Each coin can be used unlimited times.

## Approach

Use unbounded knapsack DP. Iterate over coins in the outer loop (to avoid counting permutations) and amounts in the inner loop. dp[j] += dp[j-coin] accumulates the number of ways to form amount j using coins considered so far.
