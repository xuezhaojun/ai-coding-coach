# Letter Combinations of a Phone Number

- **Difficulty**: Medium
- **Category**: Backtracking
- **Topics**: backtracking, string, hash map

## Description

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent, based on the phone keypad mapping.

## Approach

Map each digit to its corresponding letters. Use backtracking to build combinations character by character. At each digit position, iterate through all mapped letters, append each to the current combination, recurse to the next digit, then backtrack by removing the last character.
