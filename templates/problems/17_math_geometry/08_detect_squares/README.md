# Detect Squares

- **Difficulty**: Medium
- **Category**: Math & Geometry
- **Topics**: array, hash table, design, counting

## Description

Design a data structure that supports adding points on a 2D plane and counting the number of axis-aligned squares that can be formed with a given query point.

## Approach

Store all points and a frequency map. For each query point, iterate over all stored points as potential diagonal corners. If the x and y differences have equal absolute values (forming a square), check for the other two corners using the frequency map and multiply their counts.
