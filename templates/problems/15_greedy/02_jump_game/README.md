# Jump Game

- **Difficulty**: Medium
- **Category**: Greedy
- **Topics**: array, greedy

## Description

Given an integer array nums where each element represents the maximum jump length from that position, determine if you can reach the last index starting from the first index.

## Approach

Greedily track the farthest reachable index. Iterate through the array and update the maximum reachable position. If the current index ever exceeds the maximum reachable position, return false.
