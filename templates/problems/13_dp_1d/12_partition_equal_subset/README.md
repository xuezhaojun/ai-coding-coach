# Partition Equal Subset Sum

- **Difficulty**: Medium
- **Category**: 1D Dynamic Programming
- **Topics**: dynamic programming, 0/1 knapsack

## Description

Given an integer array, determine if it can be partitioned into two subsets with equal sum.

## Approach

This is a 0/1 knapsack problem. If the total sum is odd, return false. Otherwise, find if a subset sums to total/2. Use a 1D boolean DP array and iterate backwards through possible sums for each number to avoid using the same element twice.
