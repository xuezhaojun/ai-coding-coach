# 3Sum

- **Difficulty**: Medium
- **Category**: Two Pointers
- **Topics**: array, two pointers, sorting

## Description

Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, j != k, and nums[i] + nums[j] + nums[k] == 0. The solution set must not contain duplicate triplets.

## Approach

Sort the array, then fix one element and use two pointers on the remaining portion to find pairs that sum to the negative of the fixed element. Skip duplicate values at each level to avoid duplicate triplets.
