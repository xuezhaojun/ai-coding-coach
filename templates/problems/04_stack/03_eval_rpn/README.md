# Evaluate Reverse Polish Notation

- **Difficulty**: Medium
- **Category**: Stack
- **Topics**: stack, math, string

## Description

Evaluate the value of an arithmetic expression in Reverse Polish Notation. Valid operators are +, -, *, and /. Each operand may be an integer or another expression. Division truncates toward zero.

## Approach

Use a stack to process tokens left to right. Push numbers onto the stack. When an operator is encountered, pop two operands, apply the operator, and push the result back.
