# Sliding Window Maximum

- **Difficulty**: Hard
- **Category**: Sliding Window
- **Topics**: array, sliding window, monotonic deque

## Description

Given an array of integers nums and a sliding window of size k that moves from left to right, return the maximum value in each window position.

## Approach

Use a monotonic decreasing deque that stores indices. For each new element, remove all smaller elements from the back of the deque (they can never be the maximum), and remove the front if it falls outside the window. The front of the deque always holds the index of the current window's maximum.
