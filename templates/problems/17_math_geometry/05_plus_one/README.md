# Plus One

- **Difficulty**: Easy
- **Category**: Math & Geometry
- **Topics**: array, math

## Description

Given a large integer represented as an integer array digits, increment the number by one and return the resulting array of digits.

## Approach

Traverse from the last digit backward. If a digit is less than 9, increment and return. Otherwise set it to 0 and continue (carry). If all digits are 9, prepend a 1.
