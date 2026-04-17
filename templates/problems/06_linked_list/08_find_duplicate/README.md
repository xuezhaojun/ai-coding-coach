# Find the Duplicate Number

- **Difficulty**: Medium
- **Category**: Linked List
- **Topics**: Floyd's cycle detection, two pointers, array

## Description

Given an array of n+1 integers where each integer is between 1 and n inclusive, find the one duplicate number. Must not modify the array and use only constant extra space.

## Approach

Treat the array as a linked list where index i points to nums[i]. The duplicate creates a cycle. Use Floyd's tortoise and hare algorithm to find the cycle entrance, which is the duplicate number.
