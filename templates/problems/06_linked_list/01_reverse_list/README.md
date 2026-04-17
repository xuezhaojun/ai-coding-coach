# Reverse Linked List

- **Difficulty**: Easy
- **Category**: Linked List
- **Topics**: linked list, iteration

## Description

Given the head of a singly linked list, reverse the list and return the reversed list.

## Approach

Iterate through the list maintaining a previous pointer. For each node, save next, point current to previous, then advance both pointers. Return previous when done.
