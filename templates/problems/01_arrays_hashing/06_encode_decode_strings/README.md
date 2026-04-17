# Encode and Decode Strings

- **Difficulty**: Medium
- **Category**: Arrays & Hashing
- **Topics**: string, design

## Description

Design an algorithm to encode a list of strings to a single string and decode it back. The encoded string should be transmittable over a network and decodable back to the original list of strings.

## Approach

Use a length-prefixed encoding: for each string, write its length followed by a '#' delimiter and then the string itself (e.g., "5#hello3#bye"). To decode, read the number before '#' to know how many characters to consume next.
