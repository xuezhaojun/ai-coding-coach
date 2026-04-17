# Top K Frequent Elements

- **Difficulty**: Medium
- **Category**: Arrays & Hashing
- **Topics**: array, hash table, bucket sort, heap

## Description

Given an integer array nums and an integer k, return the k most frequent elements. The answer may be returned in any order.

## Approach

Count frequencies with a hash map, then use bucket sort where the bucket index represents the frequency. Iterate from the highest bucket downward to collect the top k elements. This achieves O(n) time, better than the O(n log k) heap approach.
