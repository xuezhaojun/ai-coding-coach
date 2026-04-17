# Container With Most Water

- **Difficulty**: Medium
- **Category**: Two Pointers
- **Topics**: array, two pointers, greedy

## Description

Given n non-negative integers representing the heights of vertical lines, find two lines that together with the x-axis form a container that holds the most water.

## Approach

Use two pointers at both ends. Calculate the area, then move the pointer with the shorter height inward, since moving the taller one can only decrease width without possibly increasing the limiting height. This greedy strategy finds the optimal answer in O(n) time.
