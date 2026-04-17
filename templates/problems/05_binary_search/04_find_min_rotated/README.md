# Find Minimum in Rotated Sorted Array

- **Difficulty**: Medium
- **Category**: Binary Search
- **Topics**: binary search, array

## Description

Given a sorted array that has been rotated between 1 and n times, find the minimum element. The array contains unique elements.

## Approach

Binary search comparing the mid element with the rightmost element. If mid > right, the minimum is in the right half; otherwise it is in the left half (including mid). Converge until lo == hi.
