# Reverse Nodes in K-Group

- **Difficulty**: Hard
- **Category**: Linked List
- **Topics**: linked list, recursion, iteration

## Description

Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list. Nodes that are left over (fewer than k) stay in their original order.

## Approach

Use a dummy node and process groups iteratively. For each group, first check if k nodes exist. If so, reverse them in place by adjusting pointers, then reconnect the reversed group with the previous part. Move the group pointer forward and repeat.
