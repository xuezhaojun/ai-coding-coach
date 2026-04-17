# Subsets

- **Difficulty**: Medium
- **Category**: Backtracking
- **Topics**: backtracking, bit manipulation, recursion

## Description

Given an integer array nums of unique elements, return all possible subsets (the power set). The solution set must not contain duplicate subsets and can be returned in any order.

## Approach

Use backtracking to explore including or excluding each element. At each recursive call, add a copy of the current subset to the result, then iterate from the current index forward, adding each element and recursing deeper before removing it (backtrack).
