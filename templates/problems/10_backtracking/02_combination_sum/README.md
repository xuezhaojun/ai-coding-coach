# Combination Sum

- **Difficulty**: Medium
- **Category**: Backtracking
- **Topics**: backtracking, recursion, sorting

## Description

Given an array of distinct integers candidates and a target integer, return all unique combinations of candidates where the chosen numbers sum to target. The same number may be chosen an unlimited number of times.

## Approach

Sort candidates first, then use backtracking. At each step, try adding each candidate from the current index onward (allowing reuse by not advancing the index). Prune branches early when the current candidate exceeds the remaining target.
