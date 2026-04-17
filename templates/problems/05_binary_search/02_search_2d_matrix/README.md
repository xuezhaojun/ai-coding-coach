# Search a 2D Matrix

- **Difficulty**: Medium
- **Category**: Binary Search
- **Topics**: binary search, matrix

## Description

Write an efficient algorithm that searches for a value in an m x n matrix. Integers in each row are sorted, and the first integer of each row is greater than the last integer of the previous row.

## Approach

Treat the matrix as a virtual flat sorted array of size m*n. Use standard binary search with index mapping: row = mid/n, col = mid%n. This achieves O(log(m*n)) time.
