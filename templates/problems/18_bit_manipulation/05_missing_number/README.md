# Missing Number

- **Difficulty**: Easy
- **Category**: Bit Manipulation
- **Topics**: array, bit manipulation, math, XOR

## Description

Given an array containing n distinct numbers in the range [0, n], return the one number that is missing from the array.

## Approach

XOR all indices 0..n with all values in the array. Since XOR is self-inverse, all paired values cancel out, leaving only the missing number.
