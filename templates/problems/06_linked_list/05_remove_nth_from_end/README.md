# Remove Nth Node From End of List

- **Difficulty**: Medium
- **Category**: Linked List
- **Topics**: linked list, two pointers

## Description

Given the head of a linked list, remove the nth node from the end of the list and return its head.

## Approach

Use two pointers with a gap of n+1 nodes. Advance the fast pointer n+1 steps ahead, then move both until fast reaches nil. The slow pointer will be right before the node to remove. Use a dummy node to handle edge cases like removing the head.
