# Best Time to Buy and Sell Stock

- **Difficulty**: Easy
- **Category**: Sliding Window
- **Topics**: array, dynamic programming

## Description

Given an array prices where prices[i] is the price of a given stock on the ith day, find the maximum profit from one transaction (buy one day, sell on a later day). If no profit is possible, return 0.

## Approach

Track the minimum price seen so far while iterating. At each price, compute the profit if selling at that price and update the maximum profit. This single-pass approach runs in O(n) time with O(1) space.
