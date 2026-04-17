# Longest Repeating Character Replacement

- **Difficulty**: Medium
- **Category**: Sliding Window
- **Topics**: string, hash table, sliding window

## Description

Given a string s and an integer k, you can choose any character and change it to any other uppercase English letter at most k times. Return the length of the longest substring containing the same letter after performing at most k changes.

## Approach

Use a sliding window tracking character frequencies. The key insight is that if (window size - max frequency in window) exceeds k, the window is invalid and must shrink from the left. The maxFreq value does not need to be decremented precisely because we only care about finding longer valid windows.
