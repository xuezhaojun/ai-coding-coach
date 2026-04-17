# Add Two Numbers

- **Difficulty**: Medium
- **Category**: Linked List
- **Topics**: linked list, math, simulation

## Description

Two non-negative integers are represented as linked lists where each node contains a single digit stored in reverse order. Add the two numbers and return the sum as a linked list.

## Approach

Traverse both lists simultaneously, adding corresponding digits plus carry. Create new nodes for each digit of the result. Continue until both lists are exhausted and carry is zero.
