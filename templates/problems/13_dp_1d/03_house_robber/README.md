# House Robber

- **Difficulty**: Medium
- **Category**: 1D Dynamic Programming
- **Topics**: dynamic programming, array

## Description

Given an array representing the amount of money in each house along a street, return the maximum amount you can rob without robbing two adjacent houses.

## Approach

Use two variables to track the maximum profit up to the previous two houses. At each house, decide whether to rob it (prev + current value) or skip it (curr). This reduces space from O(n) to O(1).
