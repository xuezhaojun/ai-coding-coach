# Non-Overlapping Intervals

- **Difficulty**: Medium
- **Category**: Intervals
- **Topics**: array, greedy, sorting, intervals

## Description

Given an array of intervals, return the minimum number of intervals you need to remove to make the rest non-overlapping.

## Approach

Sort by end time (greedy interval scheduling). Keep the interval with the earliest end time and remove any that overlap with it. Count the number of removals needed.
