# Climbing Stairs

- **Difficulty**: Easy
- **Category**: 1D Dynamic Programming
- **Topics**: dynamic programming, fibonacci

## Description

Given n stairs, you can climb 1 or 2 steps at a time. Return the number of distinct ways to reach the top.

## Approach

This is essentially the Fibonacci sequence. Use two variables to track the previous two values and iterate from 3 to n, updating them in place. This avoids the need for an array and achieves O(1) space.
