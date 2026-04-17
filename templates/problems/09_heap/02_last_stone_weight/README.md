# Last Stone Weight

- **Difficulty**: Easy
- **Category**: Heap
- **Topics**: max-heap, priority queue, simulation

## Description

Given an array of stone weights, repeatedly pick the two heaviest stones and smash them together. If they have different weights, the lighter one is destroyed and the heavier loses weight equal to the lighter. Return the weight of the last remaining stone, or 0 if none remain.

## Approach

Use a max-heap to efficiently extract the two heaviest stones each round. If they differ, push the difference back. Continue until one or zero stones remain.
