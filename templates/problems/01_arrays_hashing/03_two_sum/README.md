# Two Sum

- **Difficulty**: Easy
- **Category**: Arrays & Hashing
- **Topics**: array, hash table

## Description

Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target. Each input has exactly one solution, and you may not use the same element twice.

## Approach

Use a hash map to store each number's index as you iterate. For each element, check if (target - current) exists in the map. If found, return both indices immediately.
