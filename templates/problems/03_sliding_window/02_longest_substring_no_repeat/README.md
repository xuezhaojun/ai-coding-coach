# Longest Substring Without Repeating Characters

- **Difficulty**: Medium
- **Category**: Sliding Window
- **Topics**: string, hash table, sliding window

## Description

Given a string s, find the length of the longest substring without repeating characters.

## Approach

Use a sliding window with a hash map storing the last-seen index of each character. When a duplicate is found within the window, jump the left boundary past the previous occurrence. Track the maximum window size throughout.
