# Pow(x, n)

- **Difficulty**: Medium
- **Category**: Math & Geometry
- **Topics**: math, recursion, binary exponentiation

## Description

Implement pow(x, n), which calculates x raised to the power n.

## Approach

Use binary exponentiation (fast power). If n is negative, invert x and negate n. Square x and halve n each iteration; multiply the result by x when n is odd. This reduces the number of multiplications to O(log n).
