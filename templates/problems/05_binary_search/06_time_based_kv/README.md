# Time Based Key-Value Store

- **Difficulty**: Medium
- **Category**: Binary Search
- **Topics**: binary search, hash map, design

## Description

Design a time-based key-value store that can store multiple values for the same key at different timestamps, and retrieve the value at a given key and timestamp (returning the value with the largest timestamp less than or equal to the given one).

## Approach

Use a hash map where each key maps to a sorted list of (timestamp, value) pairs. Since timestamps are set in increasing order, the list is naturally sorted. Use binary search on Get to find the largest timestamp that does not exceed the query timestamp.
