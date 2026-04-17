# Meeting Rooms II

- **Difficulty**: Medium
- **Category**: Intervals
- **Topics**: array, sorting, heap, intervals, sweep line

## Description

Given an array of meeting time intervals, find the minimum number of conference rooms required.

## Approach

Separate start and end times into two sorted arrays. Use two pointers to sweep through events: when a start comes before the next end, a new room is needed; when an end comes first, a room is freed. Track the maximum concurrent rooms.
