# Merge Two Sorted Lists

- **Difficulty**: Easy
- **Category**: Linked List
- **Topics**: linked list, recursion, iteration

## Description

Merge two sorted linked lists and return it as a sorted list. The list should be made by splicing together the nodes of the two input lists.

## Approach

Use a dummy head node and iterate through both lists, always appending the smaller node. When one list is exhausted, attach the remainder of the other list.
