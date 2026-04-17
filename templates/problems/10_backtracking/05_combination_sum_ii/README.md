# Combination Sum II

- **Difficulty**: Medium
- **Category**: Backtracking
- **Topics**: backtracking, sorting, deduplication

## Description

Given a collection of candidate numbers (which may contain duplicates) and a target number, find all unique combinations where the candidate numbers sum to target. Each number in candidates may only be used once.

## Approach

Sort the candidates to group duplicates. Use backtracking with index advancement (i+1) to prevent reuse. Skip duplicates at the same level (if candidates[i] == candidates[i-1] and i > start) to avoid duplicate combinations. Prune when the current candidate exceeds the remaining target.
