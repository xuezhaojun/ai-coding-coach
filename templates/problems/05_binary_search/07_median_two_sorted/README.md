# Median of Two Sorted Arrays

- **Difficulty**: Hard
- **Category**: Binary Search
- **Topics**: binary search, array, divide and conquer

## Description

Given two sorted arrays nums1 and nums2, return the median of the two sorted arrays. The overall run time complexity should be O(log(min(m,n))).

## Approach

Binary search on the partition of the shorter array. For each partition position i in nums1, compute j in nums2 such that the left halves contain exactly half the total elements. Check if the partition is correct (maxLeft <= minRight on both sides). Adjust the binary search bounds until the correct partition is found.
