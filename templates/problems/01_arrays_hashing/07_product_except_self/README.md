# Product of Array Except Self

- **Difficulty**: Medium
- **Category**: Arrays & Hashing
- **Topics**: array, prefix sum

## Description

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i]. You must solve it without using division and in O(n) time.

## Approach

Use two passes: a forward pass to compute prefix products (product of all elements before index i) and a backward pass to multiply in suffix products (product of all elements after index i). The output array doubles as the prefix product storage, achieving O(1) extra space.
