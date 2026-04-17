# Single Number

- **Difficulty**: Easy
- **Category**: Bit Manipulation
- **Topics**: array, bit manipulation, XOR

## Description

Given a non-empty array of integers where every element appears twice except for one, find that single element.

## Approach

XOR all elements together. Since a XOR a = 0 and a XOR 0 = a, all duplicate pairs cancel out, leaving only the single number.
