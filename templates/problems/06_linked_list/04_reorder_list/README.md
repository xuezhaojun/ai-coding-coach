# Reorder List

- **Difficulty**: Medium
- **Category**: Linked List
- **Topics**: linked list, two pointers, reverse

## Description

Given the head of a singly linked list L0 -> L1 -> ... -> Ln, reorder it to L0 -> Ln -> L1 -> Ln-1 -> L2 -> Ln-2 -> ... in place.

## Approach

Three steps: (1) find the middle using slow/fast pointers, (2) reverse the second half, (3) merge the two halves by interleaving nodes alternately.
