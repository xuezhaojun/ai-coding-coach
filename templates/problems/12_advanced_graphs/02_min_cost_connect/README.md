# Min Cost to Connect All Points

- **Difficulty**: Medium
- **Category**: Advanced Graphs
- **Topics**: minimum spanning tree, Prim's algorithm, heap

## Description

Given an array of points on a 2D plane, return the minimum cost to connect all points. The cost of connecting two points is the Manhattan distance between them. All points must be connected with exactly one simple path between any two points.

## Approach

Use Prim's algorithm with a min-heap. Start from point 0 and greedily add the cheapest edge connecting a visited node to an unvisited node. The Manhattan distance serves as the edge weight. Continue until all n points are included in the MST.
