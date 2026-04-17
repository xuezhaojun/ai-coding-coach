# Insert Interval

- **Difficulty**: Medium
- **Category**: Intervals
- **Topics**: array, intervals

## Description

Given a sorted list of non-overlapping intervals and a new interval, insert the new interval and merge if necessary. Return the resulting list of non-overlapping intervals.

## Approach

Three-phase linear scan: first add all intervals that end before the new one starts, then merge all overlapping intervals with the new interval by expanding its bounds, and finally append the remaining intervals.
