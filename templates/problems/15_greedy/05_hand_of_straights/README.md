# Hand of Straights

- **Difficulty**: Medium
- **Category**: Greedy
- **Topics**: array, hash table, greedy, sorting

## Description

Given an array of integers hand and a group size, determine whether the cards can be rearranged into groups of consecutive cards of the given size.

## Approach

Sort the hand and use a frequency map. For each smallest available card, try to form a group of consecutive cards starting from it. If any card in the sequence is unavailable, return false.
