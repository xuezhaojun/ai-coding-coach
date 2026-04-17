# Valid Palindrome

- **Difficulty**: Easy
- **Category**: Two Pointers
- **Topics**: string, two pointers

## Description

Given a string s, return true if it is a palindrome after converting all uppercase letters to lowercase and removing all non-alphanumeric characters.

## Approach

Use two pointers starting from both ends. Skip non-alphanumeric characters, then compare the lowercase versions of the characters at each pointer. Move inward until the pointers meet.
