# Number of 1 Bits

- **Difficulty**: Easy
- **Category**: Bit Manipulation
- **Topics**: bit manipulation

## Description

Given an unsigned integer, return the number of 1 bits (Hamming weight).

## Approach

Use the Brian Kernighan trick: n & (n-1) clears the lowest set bit. Count how many times this operation can be performed until n becomes zero.
