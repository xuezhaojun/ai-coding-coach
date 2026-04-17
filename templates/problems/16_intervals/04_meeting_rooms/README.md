# Meeting Rooms

- **Difficulty**: Easy
- **Category**: Intervals
- **Topics**: array, sorting, intervals

## Description

Given an array of meeting time intervals, determine if a person could attend all meetings (no two meetings overlap).

## Approach

Sort intervals by start time. Check if any meeting starts before the previous one ends. If so, there is a conflict and the person cannot attend all meetings.
