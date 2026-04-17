# Contains Duplicate

- **Difficulty**: Easy
- **Category**: Arrays & Hashing
- **Topics**: array, hash set

## Description

Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

## Approach

Use a hash set to track seen numbers. For each element, check if it already exists in the set; if so, return true. Otherwise, add it and continue. This gives O(n) time and O(n) space.
