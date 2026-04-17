# Network Delay Time

- **Difficulty**: Medium
- **Category**: Advanced Graphs
- **Topics**: Dijkstra's algorithm, heap, shortest path, weighted graph

## Description

Given a network of n nodes and weighted directed edges representing signal travel times, find the time it takes for a signal sent from node k to reach all nodes. Return -1 if not all nodes can be reached.

## Approach

Use Dijkstra's algorithm with a min-heap to find shortest paths from source k to all other nodes. The answer is the maximum shortest path distance among all nodes. If any node is unreachable, return -1.
