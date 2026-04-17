# Reverse Integer

- **Difficulty**: Medium
- **Category**: Bit Manipulation
- **Topics**: math

## Description

Given a signed 32-bit integer x, return x with its digits reversed. If reversing causes overflow beyond the 32-bit signed integer range, return 0.

## Approach

Extract digits from right to left using modulo and division. Before appending each digit, check if the result would overflow the 32-bit signed integer range. Handle both positive and negative numbers since Go's modulo preserves the sign.
