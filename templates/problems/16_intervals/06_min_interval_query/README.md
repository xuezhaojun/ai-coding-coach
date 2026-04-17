# Minimum Interval to Include Each Query

- **Difficulty**: Hard
- **Category**: Intervals
- **Topics**: array, sorting, heap, intervals

## Description

Given a list of intervals and a list of queries, for each query find the size of the smallest interval that contains it. Return -1 if no interval contains the query.

## Approach

Sort intervals by start time and queries by value. For each query, push all intervals starting at or before the query value onto a min-heap keyed by interval size. Pop expired intervals (end < query). The heap top gives the smallest valid interval.
