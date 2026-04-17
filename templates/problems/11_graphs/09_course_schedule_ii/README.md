# Course Schedule II

- **Difficulty**: Medium
- **Category**: Graphs
- **Topics**: topological sort, BFS, Kahn's algorithm, directed graph

## Description

Given numCourses and prerequisite pairs, return the ordering of courses you should take to finish all courses. If there are multiple valid orderings, return any. If impossible (cycle exists), return an empty array.

## Approach

Use Kahn's algorithm (BFS-based topological sort). Compute in-degrees for all nodes, enqueue nodes with in-degree 0, and process the queue by removing edges and enqueuing newly zero-in-degree nodes. If the result has fewer than numCourses nodes, a cycle exists.
