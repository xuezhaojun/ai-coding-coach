# Redundant Connection

- **Difficulty**: Medium
- **Category**: Graphs
- **Topics**: union find, graph, cycle detection

## Description

Given a graph that started as a tree with n nodes and had one additional edge added, find and return the edge that can be removed so the result is a tree. If there are multiple answers, return the one that occurs last in the input.

## Approach

Process edges one by one using Union-Find. For each edge, check if the two nodes already share the same root. If they do, this edge creates a cycle and is the redundant connection. Otherwise, union the two nodes. The last such edge found in input order is the answer.
