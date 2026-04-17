# Clone Graph

- **Difficulty**: Medium
- **Category**: Graphs
- **Topics**: DFS, BFS, hash map, graph traversal

## Description

Given a reference of a node in a connected undirected graph, return a deep copy (clone) of the graph. Each node contains a value and a list of its neighbors.

## Approach

Use DFS with a hash map mapping original nodes to their clones. For each node, create a clone, store it in the map, then recursively clone all neighbors. The map prevents infinite loops and ensures each node is cloned exactly once.
