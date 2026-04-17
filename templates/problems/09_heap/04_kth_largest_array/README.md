# Kth Largest Element in an Array

- **Difficulty**: Medium
- **Category**: Heap
- **Topics**: min-heap, priority queue, quickselect

## Description

Given an integer array nums and an integer k, return the kth largest element in the array. This is the kth largest in sorted order, not the kth distinct element.

## Approach

Use a min-heap of size k. Iterate through all elements, pushing each onto the heap and popping when size exceeds k. After processing all elements, the heap's top is the kth largest. Alternatively, quickselect achieves O(n) average time.
