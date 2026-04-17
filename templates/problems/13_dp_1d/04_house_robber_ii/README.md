# House Robber II

- **Difficulty**: Medium
- **Category**: 1D Dynamic Programming
- **Topics**: dynamic programming, array

## Description

Houses are arranged in a circle. You cannot rob two adjacent houses (and the first and last houses are adjacent). Return the maximum amount you can rob.

## Approach

Since the first and last houses are adjacent, we cannot rob both. Run the House Robber I algorithm twice: once on houses [0..n-2] and once on [1..n-1], and return the maximum of the two results.
