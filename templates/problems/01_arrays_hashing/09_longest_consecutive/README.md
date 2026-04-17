# Longest Consecutive Sequence

- **Difficulty**: Medium
- **Category**: Arrays & Hashing
- **Topics**: array, hash table, union find

## Description

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence. The algorithm must run in O(n) time.

## Approach

Insert all numbers into a hash set. For each number, check if it is the start of a sequence (i.e., num-1 is not in the set). If so, count consecutive numbers upward. Track the maximum length found. Each element is visited at most twice, giving O(n) time.
