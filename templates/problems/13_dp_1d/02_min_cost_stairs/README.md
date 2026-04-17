# Min Cost Climbing Stairs

- **Difficulty**: Easy
- **Category**: 1D Dynamic Programming
- **Topics**: dynamic programming, array

## Description

Given an array of costs where cost[i] is the cost of stepping on the i-th stair, find the minimum cost to reach the top of the floor. You can start from step 0 or step 1, and at each step you can climb 1 or 2 stairs.

## Approach

Use bottom-up DP with two variables tracking the minimum cost to reach the previous two steps. At each step, the cost is the current step's cost plus the minimum of the two previous costs. The answer is the minimum of the last two values.
