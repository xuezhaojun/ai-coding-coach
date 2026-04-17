# Happy Number

- **Difficulty**: Easy
- **Category**: Math & Geometry
- **Topics**: math, hash table, two pointers

## Description

A happy number is defined by repeatedly replacing the number with the sum of squares of its digits until it equals 1, or loops endlessly in a cycle. Determine if a number is happy.

## Approach

Use Floyd's cycle detection (fast/slow pointers) on the digit-square-sum sequence. If the fast pointer reaches 1, the number is happy. If fast meets slow, there is a cycle and the number is not happy.
