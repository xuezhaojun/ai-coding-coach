# Car Fleet

- **Difficulty**: Medium
- **Category**: Stack
- **Topics**: stack, sorting, monotonic stack

## Description

There are n cars going to a target. Each car has a position and speed. A car fleet is formed when a faster car catches up to a slower one. Return the number of car fleets that arrive at the target.

## Approach

Sort cars by position in descending order (closest to target first). Calculate each car's arrival time. Iterate through and count a new fleet whenever a car's arrival time exceeds the previous fleet's time, since it cannot be caught.
