# Task Scheduler

- **Difficulty**: Medium
- **Category**: Heap
- **Topics**: max-heap, greedy, scheduling

## Description

Given a list of tasks represented by characters and a cooldown interval n, find the minimum number of intervals the CPU needs to finish all tasks. The same task must be separated by at least n intervals.

## Approach

Use a max-heap of task frequencies. In each round of (n+1) slots, greedily pick the most frequent tasks. After each round, re-add tasks with remaining count. If the heap is not empty after filling a round, add idle slots for the unfilled portion.
