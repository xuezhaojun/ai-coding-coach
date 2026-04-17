# Daily Temperatures

- **Difficulty**: Medium
- **Category**: Stack
- **Topics**: stack, monotonic stack, array

## Description

Given an array of daily temperatures, return an array where each element tells how many days you have to wait until a warmer temperature. If no future warmer day exists, put 0.

## Approach

Use a monotonic decreasing stack that stores indices. For each temperature, pop indices from the stack while the current temperature is warmer, computing the day difference. This processes each element at most twice (push and pop).
