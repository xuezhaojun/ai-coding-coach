# Number of Connected Components

- **Difficulty**: Medium
- **Category**: Graphs
- **Topics**: union find, graph, connected components

## Description

Given n nodes labeled 0 to n-1 and a list of undirected edges, return the number of connected components in the graph.

## Approach

Use Union-Find with path compression and union by rank. Start with n components. For each edge, find the roots of both endpoints. If they differ, union them and decrement the component count. The final count is the number of connected components.
