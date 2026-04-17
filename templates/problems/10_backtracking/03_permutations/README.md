# Permutations

- **Difficulty**: Medium
- **Category**: Backtracking
- **Topics**: backtracking, recursion, permutation

## Description

Given an array nums of distinct integers, return all possible permutations in any order.

## Approach

Use backtracking with a boolean used array. At each position, try every unused element, mark it as used, recurse to fill the next position, then unmark it. When the current permutation reaches full length, add a copy to results.
