# Graph Valid Tree

- **Difficulty**: Medium
- **Category**: Graphs
- **Topics**: union find, graph, tree

## Description

Given n nodes labeled 0 to n-1 and a list of undirected edges, determine if these edges form a valid tree. A valid tree must be connected and have no cycles.

## Approach

A tree with n nodes has exactly n-1 edges. First check this condition. Then use Union-Find to detect cycles: for each edge, union the two endpoints. If they already share the same root, a cycle exists. If all edges pass without a cycle and edge count is n-1, it is a valid tree.
