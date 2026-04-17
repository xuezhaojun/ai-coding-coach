# K Closest Points to Origin

- **Difficulty**: Medium
- **Category**: Heap
- **Topics**: max-heap, priority queue, geometry

## Description

Given an array of points on the X-Y plane and an integer k, return the k closest points to the origin (0, 0). Distance is measured as Euclidean distance (squared comparison suffices).

## Approach

Use a max-heap of size k keyed by distance. For each point, push it onto the heap and pop if size exceeds k. The heap retains the k closest points since the farthest among them is always popped.
