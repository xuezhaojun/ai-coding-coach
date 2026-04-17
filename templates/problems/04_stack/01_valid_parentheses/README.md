# Valid Parentheses

- **Difficulty**: Easy
- **Category**: Stack
- **Topics**: stack, string, hash map

## Description

Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid. A string is valid if open brackets are closed by the same type of brackets in the correct order.

## Approach

Use a stack to track opening brackets. For each closing bracket, check if it matches the top of the stack. The string is valid if the stack is empty at the end.
