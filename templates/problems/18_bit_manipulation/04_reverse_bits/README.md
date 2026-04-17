# Reverse Bits

- **Difficulty**: Easy
- **Category**: Bit Manipulation
- **Topics**: bit manipulation

## Description

Reverse the bits of a given 32-bit unsigned integer.

## Approach

Iterate 32 times: shift the result left by 1, OR with the lowest bit of the input, then shift the input right by 1. This builds the reversed number bit by bit.
