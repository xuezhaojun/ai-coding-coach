# Trapping Rain Water

- **Difficulty**: Hard
- **Category**: Two Pointers
- **Topics**: array, two pointers, dynamic programming, stack

## Description

Given n non-negative integers representing an elevation map where the width of each bar is 1, compute how much water it can trap after raining.

## Approach

Use two pointers from both ends, tracking the maximum height seen from each side. Always process the side with the smaller maximum, since the water level at that position is determined by the smaller of the two maxima. Add the difference between the current max and the current height to the total water.
