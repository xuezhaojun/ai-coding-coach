# Linked List Cycle

- **Difficulty**: Easy
- **Category**: Linked List
- **Topics**: linked list, two pointers, Floyd's algorithm

## Description

Given the head of a linked list, determine if the linked list has a cycle in it.

## Approach

Use Floyd's tortoise and hare algorithm: move a slow pointer one step and a fast pointer two steps at a time. If they meet, there is a cycle. If fast reaches nil, there is no cycle.
