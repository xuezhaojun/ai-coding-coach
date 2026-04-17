# Merge K Sorted Lists

- **Difficulty**: Hard
- **Category**: Linked List
- **Topics**: linked list, heap, divide and conquer

## Description

Merge k sorted linked lists and return it as one sorted list.

## Approach

Use a min-heap of size k. Initialize by pushing the head of each list. Repeatedly pop the smallest node, append it to the result, and push its next node if it exists. This processes all N nodes with O(log k) heap operations each.
