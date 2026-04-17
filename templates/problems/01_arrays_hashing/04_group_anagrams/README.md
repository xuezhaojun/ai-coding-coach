# Group Anagrams

- **Difficulty**: Medium
- **Category**: Arrays & Hashing
- **Topics**: string, hash table, sorting

## Description

Given an array of strings strs, group the anagrams together. You can return the answer in any order. An anagram is a word formed by rearranging the letters of another word.

## Approach

Use a character-frequency array [26]byte as the hash map key for each string. Strings with identical frequency arrays are anagrams and get grouped together. This avoids sorting each string and runs in O(n * k) time.
