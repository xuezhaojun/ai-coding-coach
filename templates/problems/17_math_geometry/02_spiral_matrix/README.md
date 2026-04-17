# Spiral Matrix

- **Difficulty**: Medium
- **Category**: Math & Geometry
- **Topics**: array, matrix, simulation

## Description

Given an m x n matrix, return all elements of the matrix in spiral order.

## Approach

Maintain four boundaries (top, bottom, left, right) and traverse in the order: right along top row, down along right column, left along bottom row, up along left column. Shrink boundaries after each direction and check bounds to avoid double-counting.
