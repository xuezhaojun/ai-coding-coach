# Kth Largest Element in a Stream

- **Difficulty**: Easy
- **Category**: Heap
- **Topics**: min-heap, priority queue, streaming

## Description

Design a class to find the kth largest element in a stream. The class is initialized with an integer k and an array of integers. Each call to add returns the kth largest element after inserting the new value.

## Approach

Maintain a min-heap of size k. The top of the heap is always the kth largest element. When adding a new element, push it onto the heap and pop if size exceeds k.
