# Search in Rotated Sorted Array

- **Difficulty**: Medium
- **Category**: Binary Search
- **Topics**: binary search, array

## Description

Given a rotated sorted array with unique values and a target, return the index of target or -1 if not found. Must achieve O(log n) runtime.

## Approach

Modified binary search: at each step determine which half is sorted by comparing nums[lo] with nums[mid]. Check if the target falls within the sorted half's range, then narrow the search accordingly.
