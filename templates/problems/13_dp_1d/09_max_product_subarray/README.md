# Maximum Product Subarray

- **Difficulty**: Medium
- **Category**: 1D Dynamic Programming
- **Topics**: dynamic programming, array

## Description

Given an integer array, find the contiguous subarray with the largest product and return the product.

## Approach

Track both the current maximum and current minimum product ending at each position, because a negative number can turn the smallest product into the largest. When encountering a negative number, swap max and min before updating. This handles all sign combinations in O(n) time.
