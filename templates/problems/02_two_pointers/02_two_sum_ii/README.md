# Two Sum II - Input Array Is Sorted

- **Difficulty**: Medium
- **Category**: Two Pointers
- **Topics**: array, two pointers, binary search

## Description

Given a 1-indexed sorted array of integers, find two numbers that add up to a specific target. Return the indices of the two numbers (1-indexed).

## Approach

Use two pointers at the start and end of the array. If the sum is too small, move the left pointer right; if too large, move the right pointer left. The sorted property guarantees convergence to the answer in O(n) time with O(1) space.
