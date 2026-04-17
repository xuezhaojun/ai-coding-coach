# Multiply Strings

- **Difficulty**: Medium
- **Category**: Math & Geometry
- **Topics**: math, string, simulation

## Description

Given two non-negative integers represented as strings, return the product as a string. You must not use any built-in big integer library or convert the inputs directly to integers.

## Approach

Use grade-school multiplication: for each pair of digits, compute the product and add it to the correct position in a result array. The product of digit at index i and j contributes to positions i+j and i+j+1. Build the final string by skipping leading zeros.
