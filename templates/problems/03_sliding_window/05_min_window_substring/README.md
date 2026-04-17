# Minimum Window Substring

- **Difficulty**: Hard
- **Category**: Sliding Window
- **Topics**: string, hash table, sliding window

## Description

Given two strings s and t, return the minimum window substring of s such that every character in t (including duplicates) is included in the window. If no such substring exists, return the empty string.

## Approach

Use a sliding window with two frequency arrays. Expand the right pointer to include characters, tracking how many distinct required characters are fully satisfied. Once all are satisfied, shrink from the left to find the minimum valid window. Track the best window found throughout the process.
