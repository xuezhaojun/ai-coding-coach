# Generate Parentheses

- **Difficulty**: Medium
- **Category**: Stack
- **Topics**: backtracking, string, recursion

## Description

Given n pairs of parentheses, generate all combinations of well-formed parentheses.

## Approach

Use backtracking to build strings character by character. At each step, add an opening parenthesis if fewer than n have been used, or a closing parenthesis if fewer closing than opening have been placed. This ensures only valid combinations are generated.
