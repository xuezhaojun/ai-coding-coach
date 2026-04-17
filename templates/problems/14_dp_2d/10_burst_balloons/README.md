# Burst Balloons

- **Difficulty**: Hard
- **Category**: 2D Dynamic Programming
- **Topics**: dynamic programming, interval DP

## Description

Given n balloons with values, bursting balloon i earns nums[i-1] * nums[i] * nums[i+1] coins. Find the maximum coins collectible by bursting all balloons.

## Approach

Use interval DP. Add boundary balloons with value 1. For each interval [left, right], try each balloon k as the LAST one to burst. The coins earned are vals[left-1] * vals[k] * vals[right+1] plus the results of the two sub-intervals. Build up from smaller to larger intervals.
