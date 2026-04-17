# Maximum Subarray

- **Difficulty**: Medium
- **Category**: Greedy
- **Topics**: array, dynamic programming, greedy, Kadane's algorithm

## Description

Given an integer array nums, find the subarray with the largest sum, and return its sum.

## Approach

Use Kadane's algorithm: maintain a running sum and reset it when it drops below zero. Track the maximum sum seen so far. This works because a negative prefix sum can never improve any future subarray.
