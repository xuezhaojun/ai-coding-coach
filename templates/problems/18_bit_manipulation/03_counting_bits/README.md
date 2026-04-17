# Counting Bits

- **Difficulty**: Easy
- **Category**: Bit Manipulation
- **Topics**: bit manipulation, dynamic programming

## Description

Given an integer n, return an array of length n+1 where element i contains the number of 1 bits in the binary representation of i.

## Approach

Use dynamic programming: the number of 1 bits in i equals the number of 1 bits in i>>1 (right-shifted by 1) plus the lowest bit (i & 1). This gives an O(n) solution without counting bits individually.
