# Largest Rectangle in Histogram

- **Difficulty**: Hard
- **Category**: Stack
- **Topics**: stack, monotonic stack, array

## Description

Given an array of integers representing the histogram's bar heights where each bar has width 1, find the area of the largest rectangle in the histogram.

## Approach

Use a monotonic increasing stack of indices. When a shorter bar is encountered, pop taller bars and calculate their maximum rectangle width (bounded by current index and new stack top). Process a virtual bar of height 0 at the end to flush remaining elements.
