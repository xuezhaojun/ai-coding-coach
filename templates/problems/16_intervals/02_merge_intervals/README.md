# Merge Intervals

- **Difficulty**: Medium
- **Category**: Intervals
- **Topics**: array, sorting, intervals

## Description

Given an array of intervals where intervals[i] = [start, end], merge all overlapping intervals and return an array of the non-overlapping intervals.

## Approach

Sort intervals by start time. Iterate through and merge each interval with the last one in the result if they overlap (current start <= previous end). Otherwise, append as a new interval.
