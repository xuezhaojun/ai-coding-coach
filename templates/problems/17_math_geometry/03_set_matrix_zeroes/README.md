# Set Matrix Zeroes

- **Difficulty**: Medium
- **Category**: Math & Geometry
- **Topics**: array, matrix, hash table

## Description

Given an m x n integer matrix, if an element is 0, set its entire row and column to 0. Do it in-place.

## Approach

Use the first row and first column as markers to record which rows and columns need to be zeroed. Use two boolean flags to track whether the first row/column themselves should be zeroed. This achieves O(1) extra space.
