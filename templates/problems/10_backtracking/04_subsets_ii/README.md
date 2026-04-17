# Subsets II

- **Difficulty**: Medium
- **Category**: Backtracking
- **Topics**: backtracking, sorting, deduplication

## Description

Given an integer array nums that may contain duplicates, return all possible subsets (the power set). The solution set must not contain duplicate subsets.

## Approach

Sort the array first to group duplicates together. Use backtracking similar to Subsets I, but skip duplicate elements at the same recursion level (if nums[i] == nums[i-1] and i > start, skip). This ensures each unique subset is generated exactly once.
