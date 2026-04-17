# Sum of Two Integers

- **Difficulty**: Medium
- **Category**: Bit Manipulation
- **Topics**: bit manipulation, math

## Description

Calculate the sum of two integers without using the + or - operators.

## Approach

Use XOR for addition without carry and AND followed by left shift for the carry. Repeat until there is no carry left. XOR gives the sum bits, AND gives the carry bits, and shifting the carry left propagates it to the correct position.
