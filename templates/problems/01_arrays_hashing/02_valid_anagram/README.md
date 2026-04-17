# Valid Anagram

- **Difficulty**: Easy
- **Category**: Arrays & Hashing
- **Topics**: string, hash table, sorting

## Description

Given two strings s and t, return true if t is an anagram of s, and false otherwise. An anagram is a word formed by rearranging the letters of another word, using all the original letters exactly once.

## Approach

Count character frequencies using a fixed-size array of 26 letters. Increment for characters in s and decrement for characters in t. If all counts are zero at the end, the strings are anagrams.
