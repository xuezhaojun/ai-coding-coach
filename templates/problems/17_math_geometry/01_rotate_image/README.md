# Rotate Image

- **Difficulty**: Medium
- **Category**: Math & Geometry
- **Topics**: array, matrix, math

## Description

Given an n x n 2D matrix representing an image, rotate the image by 90 degrees clockwise in-place.

## Approach

Perform two operations: first transpose the matrix (swap elements across the main diagonal), then reverse each row. The composition of transpose + row-reverse equals a 90-degree clockwise rotation.
