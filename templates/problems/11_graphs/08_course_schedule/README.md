# Course Schedule

- **Difficulty**: Medium
- **Category**: Graphs
- **Topics**: DFS, topological sort, cycle detection, directed graph

## Description

There are numCourses courses labeled 0 to numCourses-1. Some courses have prerequisites. Given the total number of courses and a list of prerequisite pairs, determine if it is possible to finish all courses.

## Approach

Model courses as a directed graph. Use DFS with three-state coloring (unvisited, visiting, visited) to detect cycles. If a node in the "visiting" state is encountered during DFS, a cycle exists and it is impossible to finish all courses.
