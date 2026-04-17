# LRU Cache

- **Difficulty**: Medium
- **Category**: Linked List
- **Topics**: linked list, hash map, design, doubly linked list

## Description

Design a data structure that follows the constraints of a Least Recently Used (LRU) cache. Implement Get and Put operations, both in O(1) time complexity.

## Approach

Combine a hash map (for O(1) key lookup) with a doubly linked list (for O(1) insertion/removal). The most recently used items are at the front of the list. On access, move the node to front. On eviction, remove from the tail.
