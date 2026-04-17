# Min Stack

- **Difficulty**: Medium
- **Category**: Stack
- **Topics**: stack, design

## Description

Design a stack that supports push, pop, top, and retrieving the minimum element, all in constant time.

## Approach

Maintain two stacks: one for normal values and one for tracking minimums. Push to the min stack only when the new value is less than or equal to the current minimum. Pop from the min stack when the popped value equals the current minimum.
