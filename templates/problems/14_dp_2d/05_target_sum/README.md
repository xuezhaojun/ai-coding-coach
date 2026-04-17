# Target Sum

- **Difficulty**: Medium
- **Category**: 2D Dynamic Programming
- **Topics**: dynamic programming, 0/1 knapsack

## Description

Given an array of integers and a target, assign + or - to each integer so that the sum equals the target. Return the number of ways to do this.

## Approach

Transform into a subset sum problem: if P is the sum of positive elements and N is the sum of negatives, then P - N = target and P + N = total, so P = (total + target) / 2. Use 0/1 knapsack DP to count subsets summing to P.
