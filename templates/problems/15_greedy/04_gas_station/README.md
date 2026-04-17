# Gas Station

- **Difficulty**: Medium
- **Category**: Greedy
- **Topics**: array, greedy

## Description

There are n gas stations along a circular route. Given arrays gas and cost representing the gas available and cost to travel to the next station, return the starting station index if you can travel around the circuit once, or -1 if impossible.

## Approach

Track total surplus and current surplus. If current surplus drops below zero, the start must be after the current station. If total surplus is non-negative, a valid starting station is guaranteed to exist, and our greedy choice of start is correct.
