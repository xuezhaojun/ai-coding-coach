# Copy List with Random Pointer

- **Difficulty**: Medium
- **Category**: Linked List
- **Topics**: linked list, hash map, deep copy

## Description

A linked list where each node has an additional random pointer that could point to any node or null. Construct a deep copy of the list.

## Approach

Two passes with a hash map: first pass creates all new nodes and maps old nodes to new nodes. Second pass sets the Next and Random pointers for each new node using the map.
